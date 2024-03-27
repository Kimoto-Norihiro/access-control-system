package repository

import (
	"time"

	"gorm.io/gorm"

	"github.com/Kimoto-Norihiro/access-control-system/model"
)

type RecordRepository interface {
	// 最新の在室情報取得
	GetLatestRecord(db *gorm.DB, userID int) (*model.Record, error)
	// 入室
	Entry(tx *gorm.DB, userID int, entryAt time.Time) error
	// 退室
	Exit(tx *gorm.DB, record *model.Record, exitAt time.Time) error
}

type recordRepository struct {
}

func NewRecordRepository() RecordRepository {
	return &recordRepository{}
}

func (r *recordRepository) GetLatestRecord(db *gorm.DB, userID int) (*model.Record, error) {
	var record model.Record
	if err := db.Where("user_id = ?", userID).Last(&record).Error; err != nil {
		return nil, err
	}

	return &record, nil
}

func (r *recordRepository) Entry(tx *gorm.DB, userID int, entryAt time.Time) error {
	record := &model.Record{
		UserID:  userID,
		EntryAt: entryAt,
	}

	if err := tx.Create(record).Error; err != nil {
		return err
	}

	return nil
}

func (r *recordRepository) Exit(tx *gorm.DB, record *model.Record, exitAt time.Time) error {
	record.ExitAt = &exitAt
	if err := tx.Save(record).Error; err != nil {
		return err
	}

	return nil
}
