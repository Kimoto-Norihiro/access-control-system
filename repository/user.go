package repository

import (
	"gorm.io/gorm"

	"github.com/Kimoto-Norihiro/access-control-system/model"
)

type UserRepository interface {
	// ユーザー登録
	CreateUser(tx *gorm.DB, name string) error
	// ユーザーの月間在室時間の取得
	GetUserMonthlyAttendanceTime(db *gorm.DB, userID int) error
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) CreateUser(tx *gorm.DB, name string) error {
	user := &model.User{
		Name: name,
		IsExist: false,
	}

	err := tx.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetUserMonthlyAttendanceTime(db *gorm.DB, userID int) error {
	return nil
}
