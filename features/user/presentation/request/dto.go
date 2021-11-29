package request

import (
	"antonio/features/user"
)

type UserRequest struct {
	Name        string              `json: "name"`
	Bio         string              `json: "bio"`
	Gender      string              `json: "gender"`
	Email       string              `json: "email"`
	Password    string              `json: "password"`
}

type UserAuth struct {
	Email    string `json: "email"`
	Password string `json: "password"`
}


func (data *UserAuth) ToUserCore() user.UserCore {
	return user.UserCore{
		Email:    data.Email,
		Password: data.Password,
	}
}


func (requestData *UserRequest) ToUserCore() user.UserCore {
	return user.UserCore{
		Name:        requestData.Name,
		Bio:	     requestData.Bio,
		Gender:      requestData.Gender,
		Email:       requestData.Email,
		Password:    requestData.Password,
	}
}
