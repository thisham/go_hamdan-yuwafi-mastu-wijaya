package auth

import "middleware-api/src/models/user"

func (repo *LoginRepository) AttemptLogin(data Login) (string, error) {
	var user user.User
	foundRecord := repo.DB.Where("email = ? and password = ?", data.Email, data.Password).Find(&user)
	if err := foundRecord.Error; err != nil || foundRecord.RowsAffected == 0 {
		return "", err
	}
	return user.ID.String(), nil
}
