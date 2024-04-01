package usecase

type ListExistUsersOutput struct {
	UserNames []string `json:"user_names"`
}

func (u *usecase) ListExistUsers() (*ListExistUsersOutput, error) {
	users, err := u.recordRepo.ListExistUsers(u.db)
	if err != nil {
		return nil, err
	}

	var userNames []string
	for _, user := range *users {
		userNames = append(userNames, user.Name)
	}

	return &ListExistUsersOutput{
		UserNames: userNames,
	}, nil
}
