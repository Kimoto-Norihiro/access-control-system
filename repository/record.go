package repository

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/Kimoto-Norihiro/access-control-system/model"
)

type RecordRepository interface {
	// 最新の在室情報取得
	GetLatestRecord(db *gorm.DB, userID int) (*model.Record, error)
	// 入室
	Enter(tx *gorm.DB, userID int) (*model.Record, error)
	// 退室
	Exit(tx *gorm.DB, record *model.Record) (*model.Record, error)
	// 在室しているユーザーの情報を取得
	ListExistUsers(db *gorm.DB) (*[]model.User, error)
}

type recordRepository struct {
}

func NewRecordRepository() RecordRepository {
	return &recordRepository{}
}

func (r *recordRepository) GetLatestRecord(db *gorm.DB, userID int) (*model.Record, error) {
	var record model.Record
	if err := db.Where("user_id = ?", userID).Preload("User").Last(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &record, nil
}

func (r *recordRepository) Enter(tx *gorm.DB, userID int) (*model.Record, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	record := &model.Record{
		UserID:  userID,
		EnterAt: time.Now().In(jst),
	}

	if err := tx.Create(record).Error; err != nil {
		return nil, err
	}

	if err := tx.Preload("User").First(record, record.ID).Error; err != nil {
		return nil, err
	}

	return record, nil
}

func (r *recordRepository) Exit(tx *gorm.DB, record *model.Record) (*model.Record, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	now := time.Now().In(jst)
	record.ExitAt = &now
	if err := tx.Save(record).Preload("User").Error; err != nil {
		return nil, err
	}

	return record, nil
}

func (r *recordRepository) ListExistUsers(db *gorm.DB) (*[]model.User, error) {
	var records []model.Record
	// 入室順で取得
	if err := db.Where("exit_at IS NULL").Preload("User").Order("enter_at").Find(&records).Error; err != nil {
		return nil, err
	}

	var users []model.User
	for _, record := range records {
		users = append(users, record.User)
	}

	return &users, nil
}
