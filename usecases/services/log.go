package services

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/segmentio/kafka-go"
	"github.com/vnFuhung2903/vcs-logging-services/models"
	"github.com/vnFuhung2903/vcs-logging-services/usecases/repositories"
)

type LogService interface {
	FindLogs(processed bool) ([]*models.Log, error)
	Process() error
	DeleteProcessedLogs() error
}

type logService struct {
	Lr         repositories.LogRepository
	Writer     *kafka.Writer
	NumWorkers int
	BatchSize  int
}

func NewLogService(lr repositories.LogRepository, writer *kafka.Writer, numWorkers int, batchSize int) LogService {
	return &logService{
		Lr:         lr,
		Writer:     writer,
		NumWorkers: numWorkers,
		BatchSize:  batchSize,
	}
}

func (logService *logService) FindLogs(processed bool) ([]*models.Log, error) {
	logs, err := logService.Lr.FindLogs(processed, logService.BatchSize)
	if err != nil {
		return nil, err
	}
	return logs, nil
}

func (logService *logService) Process() error {
	logs, err := logService.FindLogs(false)
	if err != nil {
		return err
	}

	jobs := make(chan []*models.Log, len(logs)/logService.BatchSize+1)
	errChan := make(chan error, logService.NumWorkers)
	var wg sync.WaitGroup

	for range logService.NumWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for batch := range jobs {
				err := logService.publishLogs(batch)
				if err != nil {
					errChan <- err
					return
				} else {
					err = logService.Lr.UpdateLogs(batch)
					if err != nil {
						errChan <- err
						return
					}
				}
			}
		}()
	}

	for i := 0; i < len(logs); i += logService.BatchSize {
		j := min(i+logService.BatchSize, len(logs))
		jobs <- logs[i:j]
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(errChan)
	}()
	for err := range errChan {
		if err != nil {
			return err
		}
	}
	return nil
}

func (logService *logService) DeleteProcessedLogs() error {
	err := logService.Lr.DeleteProcessedLogs()
	return err
}

func (logService *logService) publishLogs(logs []*models.Log) error {
	for _, log := range logs {
		msg, err := json.Marshal(log)
		if err != nil {
			return err
		}

		err = logService.Writer.WriteMessages(
			context.Background(),
			kafka.Message{
				Key:   fmt.Appendf(nil, "%d", log.Id),
				Value: msg,
			},
		)
		if err != nil {
			return err
		}
	}
	return nil
}
