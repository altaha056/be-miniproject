package data

import (
	"antonio/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       	int
	Username	string
	Email    	string
	Password 	string
	Bio			string
	Gender		string
}

func toUserRecord(user users.Core) User {
	return User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Bio: 	  user.Bio,
		Gender:   user.Gender,
	}
}

func toUserCore(user User) users.Core {
	return users.Core{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Bio: 	  user.Bio,
		Gender:   user.Gender,
	}
}

func toUserCoreList(uList []User) []users.Core {
	convertedUser := []users.Core{}

	for _, user := range uList {
		convertedUser = append(convertedUser, toUserCore(user))
	}

	return convertedUser
}

func toUserRecordList(uList []users.Core) []User {
	convertedUser := []User{}

	for _, user := range uList {
		convertedUser = append(convertedUser, toUserRecord(users.Core{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Bio: 	  user.Bio,
			Gender:   user.Gender,
		}))
	}

	return convertedUser
}
