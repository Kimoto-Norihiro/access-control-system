package usecase

import (
	"errors"

	"github.com/Kimoto-Norihiro/access-control-system/model"
	"gorm.io/gorm"
)

type EntryInput struct {
	UserID int `json:"user_id" binding:"required"`
}

type EntryOutput struct {
	UserName string `json:"user_name"`
}

func (u *usecase) Entry(input *EntryInput) (*EntryOutput, error) {
	var user *model.User
	err := u.db.Transaction(func(tx *gorm.DB) error {
		latestRecord, err := u.recordRepo.GetLatestRecord(tx, input.UserID)
		if err != nil {
			return err
		}

		if latestRecord.ExitAt == nil {
			return errors.New("already entered")
		}

		user, err = u.recordRepo.Entry(tx, input.UserID)
		if err != nil {
			return err
		}

		return nil
	})

	if user == nil {
		return nil, err
	}

	return &EntryOutput{
		UserName: user.Name,
	}, err
}
