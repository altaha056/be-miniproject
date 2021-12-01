package response

import (
	"antonio/features/users"
)

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Bio 	 string `json:"bio"`
	Gender   string `json:"gender"`
}

func ToUserResponse(user users.Core) UserResponse {
	return UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Bio: 	  user.Bio,
		Gender:   user.Gender,
	}
}

func ToUserResponseList(userList []users.Core) []UserResponse {
	convertedUser := []UserResponse{}
	for _, user := range userList {
		convertedUser = append(convertedUser, ToUserResponse(user))
	}

	return convertedUser
}
