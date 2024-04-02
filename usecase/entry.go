package usecase

import (
	"errors"

	"gorm.io/gorm"
)

type EntryInput struct {
	UserID int `json:"user_id" binding:"required"`
}

type EntryOutput struct {
	UserName string `json:"user_name"`
	EntryAt  string `json:"entry_at"`
}

func (u *usecase) Entry(input *EntryInput) (*EntryOutput, error) {
	var o *EntryOutput

	err := u.db.Transaction(func(tx *gorm.DB) error {
		// 最新の在室情報取得
		latestRecord, _ := u.recordRepo.GetLatestRecord(tx, input.UserID)

		// すでに入室している場合はエラー
		if latestRecord != nil && latestRecord.ExitAt == nil {
			return errors.New("already entered")
		}

		record, err := u.recordRepo.Entry(tx, input.UserID)
		if err != nil {
			return err
		}

		o = &EntryOutput{
			UserName: record.User.Name,
			EntryAt:  record.EntryAt.Format("2006-01-02 15:04:05"),
		}

		return nil
	})

	return o, err
}
