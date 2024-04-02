package repository

import (
	"gorm.io/gorm"

	"github.com/Kimoto-Norihiro/access-control-system/model"
)

type UserRepository interface {
	// ユーザー登録
	CreateUser(tx *gorm.DB, name string) error
	// ユーザー一覧
	ListUsers(db *gorm.DB) (*[]model.User, error)
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
	}

	err := tx.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) ListUsers(db *gorm.DB) (*[]model.User, error) {
	var users []model.User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}

func (r *userRepository) GetUserMonthlyAttendanceTime(db *gorm.DB, userID int) error {
	return nil
}
