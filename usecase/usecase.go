package usecase

import (
	"gorm.io/gorm"

	"github.com/Kimoto-Norihiro/access-control-system/repository"
)

type Usecase interface {
	// ユーザー作成
	CreateUser(input *CreateUserInput) error
	// ユーザー一覧
	ListUsers() (*ListUsersOutput, error)
	// 最新の在室情報取得
	GetLatestRecord(input *GetLatestRecordInput) (*GetLatestRecordOutput, error)
	// 入室
	Enter(input *EnterInput) (*EnterOutput, error)
	// 退室
	Exit(input *ExitInput) (*ExitOutput, error)
	// 在室しているユーザーの情報を取得
	ListExistUsers() (*ListExistUsersOutput, error)
}

type usecase struct {
	db         *gorm.DB
	userRepo   repository.UserRepository
	recordRepo repository.RecordRepository
}

func NewUsecase(db *gorm.DB, userRepo repository.UserRepository, recordRepo repository.RecordRepository) Usecase {
	return &usecase{
		db:         db,
		userRepo:   userRepo,
		recordRepo: recordRepo,
	}
}
