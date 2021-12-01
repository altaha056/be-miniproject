package business

import (
	"antonio/features/user"
	"antonio/middleware"
)

type userService struct {
	userRepository user.Repository
	
}

func NewUserService(userRepository user.Repository) user.Business {
	return &userService{userRepository}
}

func (us *userService) RegisterUser(data user.UserCore) error {
	err := us.userRepository.InsertData(data)

	if err != nil {
		return err
	}

	return nil
}

func (us *userService) GetUsers(data user.UserCore) ([]user.UserCore, error) {
	users, err := us.userRepository.GetData(data)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (us *userService) LoginUser(data user.UserCore) (user.UserCore, error) {
	userData, err := us.userRepository.CheckUser(data)
	if err != nil {
		return user.UserCore{}, err
	}

	userData.Token, err = middleware.CreateToken(userData.Id, "user")
	if err != nil {
		return user.UserCore{}, err
	}

	return userData, nil
}

func (us *userService) GetUserById(id int) (user.UserCore, error) {
	userData, err := us.userRepository.GetDataById(id)

	if err != nil {
		return user.UserCore{}, err
	}

	return userData, nil
}