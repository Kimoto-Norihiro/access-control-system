package usecase

import (
	"errors"

	"github.com/Kimoto-Norihiro/access-control-system/model"
	"gorm.io/gorm"
)

type EnterInput struct {
	UserID int `json:"user_id" binding:"required"`
}

type EnterOutput struct {
	UserName string       `json:"user_name"`
	EnterAt  string       `json:"enter_at"`
	Record   model.Record `json:"record"`
}

func (u *usecase) Enter(input *EnterInput) (*EnterOutput, error) {
	var o *EnterOutput

	err := u.db.Transaction(func(tx *gorm.DB) error {
		// 最新の在室情報取得
		latestRecord, _ := u.recordRepo.GetLatestRecord(tx, input.UserID)

		// すでに入室している場合はエラー
		if latestRecord != nil && latestRecord.ExitAt == nil {
			return errors.New("already entered")
		}

		record, err := u.recordRepo.Enter(tx, input.UserID)
		if err != nil {
			return err
		}

		o = &EnterOutput{
			UserName: record.User.Name,
			EnterAt:  record.EnterAt.Format("2006-01-02 15:04:05"),
			Record:   *record,
		}

		return nil
	})

	return o, err
}
