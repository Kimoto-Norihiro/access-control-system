package usecase

import (
	"errors"
	"time"

	"github.com/Kimoto-Norihiro/access-control-system/model"
	"gorm.io/gorm"
)

type ExitInput struct {
	UserID int       `json:"user_id" binding:"required"`
	ExitAt time.Time `json:"exit_at" binding:"required"`
}

type ExitOutput struct {
	UserName string `json:"user_name"`
}

func (u *usecase) Exit(input *ExitInput) (*ExitOutput, error) {
	var user *model.User
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
		user, err = u.recordRepo.Exit(tx, latestRecord)
		if err != nil {
			return err
		}

		return nil
	})

	if user == nil {
		return nil, err
	}

	return &ExitOutput{
		UserName: user.Name,
	}, err
}
