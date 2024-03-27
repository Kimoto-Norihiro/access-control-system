package usecase

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type ExitInput struct {
	UserID int       `json:"user_id" binding:"required"`
	ExitAt time.Time `json:"exit_at" binding:"required"`
}

func (u *usecase) Exit(input *ExitInput) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		// 最新の在室情報取得
		latestRecord, err := u.recordRepo.GetLatestRecord(tx, input.UserID)
		if err != nil {
			return err
		}

		if latestRecord == nil {
			return errors.New("record not found")
		}

		if latestRecord.ExitAt != nil {
			return errors.New("already exited")
		}

		// 退室処理
		return u.recordRepo.Exit(tx, latestRecord, input.ExitAt)
	})
	return err
}
