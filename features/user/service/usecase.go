package service

import (
	"antonio/features/user"
	"antonio/middleware"
)

type userService struct {
	userRepository user.Repository
}

func NewUserService(userRepository user.Repository) user.Service{
	return &userService{userRepository}
}

func (usv *userService) RegisterUser(data user.UserCore) error{
	err:=usv.userRepository.InsertData(data)
	if err!=nil{
		return err
	}
	return nil
}

func (usv *userService) LoginUser(data user.UserCore) (user.UserCore,error){
	userData, err:=usv.userRepository.CheckUser(data)
	if err !=nil{
		return user.UserCore{},err
	}
	userData.Token,err=middleware.CreateToken(userData.Id,"user")
	if err!=nil{
		return user.UserCore{},err
	}
	return userData,nil
}

func (usv *userService) GetUsers(data user.UserCore)([]user.UserCore, error){
	users,err:=usv.userRepository.GetData(data)
	if err!=nil{
		return nil, err
	}
	return users,nil
}

func (usv *userService) GetUserByName(name string)(user.UserCore,error){
	userData,err:=usv.userRepository.GetDataByName(name)
	if err!=nil{
		return user.UserCore{},err
	}
	return userData, nil
}