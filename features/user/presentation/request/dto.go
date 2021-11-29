package request

import "antonio/features/user"

type UserRequest struct {
	Name     string `json:"name"`
	Bio      string `json:"bio"`
	Gender   string `json:"gender"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (data *UserAuth) TouserCore() user.UserCore{
	return user.UserCore{
		Email: data.Email,
		Password: data.Password,
	}
}

func (reqData *UserRequest)TouserCore()user.UserCore{
	return user.UserCore{
		Name: reqData.Name,
		Bio: reqData.Bio,
		Gender: reqData.Gender,
		Email: reqData.Email,
		Password: reqData.Password,
	}
}