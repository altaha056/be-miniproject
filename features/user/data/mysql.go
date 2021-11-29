package data

import (
	"antonio/features/user"
	"errors"

	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	DB *gorm.DB
}

func NewMysqlUserRepository(DB *gorm.DB) user.Repository{
	return &mysqlUserRepository{DB}
}

func (mur *mysqlUserRepository) InsertData(data user.UserCore) error{
	recordData:=toUserRecord(data)
	err:=mur.DB.FirstOrCreate(&recordData)
	if err!=nil{
		return err.Error
	}
	return nil
}

func (mur *mysqlUserRepository)CheckUser(data user.UserCore)(user.UserCore, error){
	var userData User
	err:=mur.DB.Where("email=? and password=?", data.Email, data.Password).First(&userData).Error

	if userData.Name =="" && userData.ID == 0{
		return user.UserCore{}, errors.New("no user found")
	}
	if err!=nil{
		return user.UserCore{},err
	}
	return toUserCore(userData), nil
}

func (mur *mysqlUserRepository)GetDataById(id int)(user.UserCore, error){
	var userData User
	err:=mur.DB.First(&userData, id).Error

	if userData.Name=="" && userData.ID ==0{
		return user.UserCore{}, errors.New("no user found")
	}
	if err!=nil{
		return user.UserCore{}, err
	}
	return toUserCore(userData), nil
}

func (mur *mysqlUserRepository)GetData(data user.UserCore)([]user.UserCore, error){
	var users[]User

	err:=mur.DB.Debug().Distinct("users.id","users.name","users.email","users.gender").Find(&users).Error

	if err!=nil{
		return nil, err
	}

	return toUserCoreList(users), nil
}