package users

import "static-api/src/constants"

func findUserIndex(userId int) int {
	var userIndex int = -1

	for index, user := range constants.Users {
		if user.Id == userId {
			userIndex = index
			break
		}
	}

	return userIndex
}
