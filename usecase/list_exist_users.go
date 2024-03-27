package usecase

func (u *usecase) ListExistUsers() ([]string, error) {
	users, err := u.userRepo.ListExistUsers(u.db)
	if err != nil {
		return nil, err
	}

	var userNames []string
	for _, user := range *users {
		userNames = append(userNames, user.Name)
	}

	return userNames, nil
}
