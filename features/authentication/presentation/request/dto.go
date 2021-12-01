package request

import (
	"antonio/features/users"
)

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}


func (requestData *UserRequest) ToUserCore() users.Core {
	return users.Core{
		Email:    requestData.Email,
		Password: requestData.Password,
	}
}
