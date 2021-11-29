package response

import (
	"antonio/features/user"
)

type UserResponse struct {
	Id          uint                 `json: "id"`
	Name        string               `json: "name"`
	Bio         string               `json: "bio"`
	Gender      string               `json: "gender"`
	Email	    string               `json: "email"`
}


type UserLoginResponse struct {
	Id    uint   `json: "id"`
	Name  string `json: "name"`
	Email string `json: "email"`
	Token string `json: "token"`
}


func ToUserResponse(user user.UserCore) UserResponse {
	return UserResponse{
		Id:          user.Id,
		Name:        user.Name,
		Bio:         user.Bio,
		Gender:      user.Gender,
		Email:       user.Email,
	}
}

func ToUserResponseList(userList []user.UserCore) []UserResponse {
	convertedUser := []UserResponse{}
	for _, user := range userList {
		convertedUser = append(convertedUser, ToUserResponse(user))
	}

	return convertedUser
}

func ToUserLoginResponse(user user.UserCore) UserLoginResponse {
	return UserLoginResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Token: user.Token,
	}
}
