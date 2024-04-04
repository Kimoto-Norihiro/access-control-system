package usecase

import (
	"github.com/Kimoto-Norihiro/access-control-system/model"
)

type GetLatestRecordInput struct {
	UserID int `json:"user_id" binding:"required"`
}

type GetLatestRecordOutput struct {
	Record *model.Record `json:"record"`
}

func (u *usecase) GetLatestRecord(input *GetLatestRecordInput) (*GetLatestRecordOutput, error) {
	latestRecord, err := u.recordRepo.GetLatestRecord(u.db, input.UserID)
	if err != nil {
		return nil, nil
	}

	if latestRecord == nil {
		return nil, nil
	}

	if latestRecord != nil && latestRecord.ExitAt != nil {
		return nil, nil
	}

	return &GetLatestRecordOutput{
		Record: latestRecord,
	}, nil
}
