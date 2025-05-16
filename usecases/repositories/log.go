package repositories

import (
	"github.com/vnFuhung2903/vcs-logging-services/models"
	"gorm.io/gorm"
)

type LogRepository interface {
	FindAllUnprocessedLogs() ([]*models.Log, error)
	DeleteProcessedLogs() error
}

type logRepository struct {
	Db *gorm.DB
}

func NewLogRepository(db *gorm.DB) LogRepository {
	return &logRepository{Db: db}
}

func (lr *logRepository) FindAllUnprocessedLogs() ([]*models.Log, error) {
	var logs []*models.Log
	res := lr.Db.Find(&logs, models.Log{Processed: false})
	if res.Error != nil {
		return nil, res.Error
	}
	return logs, nil
}

func (lr *logRepository) DeleteProcessedLogs() error {
	res := lr.Db.Where("processed = ?", true).Delete(&models.Log{})
	return res.Error
}
