package user

func (repo *UserRepository) GetAllUsers() ([]User, error) {
	var users []User
	if err := repo.DB.Find(&users).Error; err != nil {
		return []User{}, err
	}

	return users, nil
}

func (repo *UserRepository) GetUser(id string) (User, error) {
	var user User
	if err := repo.DB.Find(&user, id).Error; err != nil {
		return User{}, err
	}

	return user, nil
}
