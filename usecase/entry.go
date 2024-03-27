package usecase

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type EntryInput struct {
	UserID  int       `json:"user_id" binding:"required"`
	EntryAt time.Time `json:"entry_at" binding:"required"`
}

func (u *usecase) Entry(input *EntryInput) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		latestRecord, err := u.recordRepo.GetLatestRecord(tx, input.UserID)
		if err != nil {
			return err
		}

		if latestRecord.ExitAt == nil {
			return errors.New("already entered")
		}

		return u.recordRepo.Entry(tx, input.UserID, input.EntryAt)
	})

	return err
}
