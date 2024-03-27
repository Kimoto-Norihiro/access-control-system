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
	Entry(tx *gorm.DB, userID int) error
	// 退室
	Exit(tx *gorm.DB, record *model.Record) error
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

func (r *recordRepository) Entry(tx *gorm.DB, userID int) error {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	record := &model.Record{
		UserID:  userID,
		EntryAt: time.Now().In(jst),
	}

	if err := tx.Create(record).Error; err != nil {
		return err
	}

	return nil
}

func (r *recordRepository) Exit(tx *gorm.DB, record *model.Record) error {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	now := time.Now().In(jst)
	record.ExitAt = &now
	if err := tx.Save(record).Error; err != nil {
		return err
	}

	return nil
}
