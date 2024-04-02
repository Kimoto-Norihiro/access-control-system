package usecase

import "github.com/Kimoto-Norihiro/access-control-system/model"

type ListUsersOutput struct {
	Users []model.User
}

func (u *usecase) ListUsers() (*ListUsersOutput, error) {
	users, err := u.userRepo.ListUsers(u.db)
	if err != nil {
		return nil, err
	}

	return &ListUsersOutput{Users: *users}, nil
}
