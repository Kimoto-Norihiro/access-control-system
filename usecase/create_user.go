package usecase

import "gorm.io/gorm"

type CreateUserInput struct {
	Name string `json:"name" binding:"required"`
}

func (u *usecase) CreateUser(input *CreateUserInput) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		return u.userRepo.CreateUser(tx, input.Name)
	})
	return err
}