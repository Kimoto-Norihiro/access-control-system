package usecase

import (
	"gorm.io/gorm"

	"github.com/Kimoto-Norihiro/access-control-system/repository"
)

type Usecase interface {
	// ユーザー作成
	CreateUser(input *CreateUserInput) error
	// 入室
	Entry(input *EntryInput) error
	// 退室
	Exit(input *ExitInput) error
	// 在室しているユーザーの情報を取得
	ListExistUsers() ([]string, error)
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