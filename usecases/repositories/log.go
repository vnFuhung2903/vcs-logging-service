package repositories

import (
	"github.com/vnFuhung2903/vcs-logging-services/models"
	"gorm.io/gorm"
)

type LogRepository interface {
	FindLogs(processed bool, limit int) ([]*models.Log, error)
	UpdateLogs(logs []*models.Log) error
	DeleteProcessedLogs() error
}

type logRepository struct {
	Db *gorm.DB
}

func NewLogRepository(db *gorm.DB) LogRepository {
	return &logRepository{Db: db}
}

func (lr *logRepository) FindLogs(processed bool, limit int) ([]*models.Log, error) {
	var logs []*models.Log
	err := lr.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Limit(limit).Find(&logs, models.Log{Processed: processed})
		return res.Error
	})
	if err != nil {
		return nil, err
	}
	return logs, nil
}

func (lr *logRepository) UpdateLogs(logs []*models.Log) error {
	var ids []uint
	for _, log := range logs {
		ids = append(ids, log.Id)
	}
	err := lr.Db.Transaction(func(tx *gorm.DB) error {
		res := lr.Db.Model(&models.Log{}).Where("id IN ?", ids).Update("processed", true)
		return res.Error
	})
	return err
}

func (lr *logRepository) DeleteProcessedLogs() error {
	res := lr.Db.Where("processed = ?", true).Delete(&models.Log{})
	return res.Error
}
