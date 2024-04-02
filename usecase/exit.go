package usecase

import (
	"errors"

	"gorm.io/gorm"
)

type ExitInput struct {
	UserID int `json:"user_id" binding:"required"`
}

type ExitOutput struct {
	UserName string `json:"user_name"`
	ExitAt   string `json:"exit_at"`
}

func (u *usecase) Exit(input *ExitInput) (*ExitOutput, error) {
	var o *ExitOutput
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

		record, err := u.recordRepo.Exit(tx, latestRecord)
		if err != nil {
			return err
		}

		o = &ExitOutput{
			UserName: record.User.Name,
			ExitAt:   record.ExitAt.Format("2006-01-02 15:04:05"),
		}

		return nil
	})

	return o, err
}
