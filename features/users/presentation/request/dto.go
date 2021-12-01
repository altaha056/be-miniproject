package request

import "antonio/features/users"

type UserRequest struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Bio 	 string `json:"bio"`
	Gender   string `json:"gender"`
}

func (requestData *UserRequest) ToUserCore() users.Core {
	return users.Core{
		ID:       requestData.ID,
		Username: requestData.Username,
		Email:    requestData.Email,
		Password: requestData.Password,
		Bio: 	  requestData.Bio,
		Gender:   requestData.Gender,
	}
}
