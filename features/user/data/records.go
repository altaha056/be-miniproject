package data

import (
	"antonio/features/user"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Bio string
	Gender string
	Email string
	Password string
}

func toUserRecord(user user.UserCore)User{
	return User{
		Name: 		user.Name,
		Bio: 		user.Bio,
		Gender:		user.Gender,
		Email: 		user.Email,
		Password: 	user.Password,
	}
}

func toUserCore(u User)user.UserCore{
	return user.UserCore{
		Id: u.ID,
		Name: u.Name,
		Bio: u.Bio,
		Gender: u.Gender,
		Email: u.Email,
		Password: u.Password,
	}
}

func toUserCoreList(ul []User) []user.UserCore{
	convertedUser:=[]user.UserCore{}

	for _,user:=range ul{
		convertedUser=append(convertedUser, toUserCore(user))
	}
	return convertedUser
}