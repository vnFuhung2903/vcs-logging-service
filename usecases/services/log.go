package services

import (
	"github.com/vnFuhung2903/vcs-logging-services/models"
	"github.com/vnFuhung2903/vcs-logging-services/usecases/repositories"
)

type LogService interface {
	FindAllUnprocessedLogs() ([]*models.Log, error)
	DeleteProcessedLogs() error
}

type logService struct {
	Lr repositories.LogRepository
}

func NewLogService(lr *repositories.LogRepository) LogService {
	return &logService{Lr: *lr}
}

func (logService *logService) FindAllUnprocessedLogs() ([]*models.Log, error) {
	logs, err := logService.Lr.FindAllUnprocessedLogs()
	if err != nil {
		return nil, err
	}
	return logs, nil
}

func (logService *logService) DeleteProcessedLogs() error {
	err := logService.Lr.DeleteProcessedLogs()
	return err
}
