package data

import (
	"antonio/features/users"

	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) users.Data {
	return &mysqlUserRepository{
		Conn: conn,
	}
}

func (ur *mysqlUserRepository) CreateUser(data users.Core) (userId int, err error) {

	recordData := toUserRecord(data)
	err = ur.Conn.Create(&recordData).Error
	if err != nil {
		return 0, err
	}
	return recordData.ID, nil
}

func (ur *mysqlUserRepository) GetAllUsers() ([]users.Core, error) {

	var users []User

	err := ur.Conn.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return toUserCoreList(users), nil

}

func (ur *mysqlUserRepository) GetUserById(userId int) (users.Core, error) {
	var user User
	err := ur.Conn.First(&user, userId).Error

	if err != nil {
		return users.Core{}, err
	}

	return toUserCore(user), nil

}
