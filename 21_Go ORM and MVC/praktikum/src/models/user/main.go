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
	if err := repo.DB.Where("ID = ?", id).Find(&user).Error; err != nil {
		return User{}, err
	}

	return user, nil
}

func (repo *UserRepository) CreateUser(user User) error {
	if err := repo.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (repo *UserRepository) UpdateUser(user User, id string) error {
	if err := repo.DB.Where("ID = ?", id).Updates(&user).Error; err != nil {
		return err
	}

	return nil
}

func (repo *UserRepository) DeleteUser(id string) error {
	if err := repo.DB.Where("ID = ?", id).Delete(&User{}).Error; err != nil {
		return err
	}

	return nil
}
