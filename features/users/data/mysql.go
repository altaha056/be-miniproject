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

// func (ur *mysqlUserRepository) GetUserFollowingById(userId int) ([]users.Core, error) {
// 	var usersFollowing []User

// 	err := ur.Conn.Raw("SELECT users.id, users.email, users.fullname, users.image FROM follows LEFT JOIN users ON follows.following_user_id = users.id WHERE follows.followers_user_id = ?", userId).Scan(&usersFollowing).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return toUserCoreList(usersFollowing), nil

// }

// func (ur *mysqlUserRepository) GetUserFollowersById(userId int) ([]users.Core, error) {

// 	var usersFollowers []User

// 	err := ur.Conn.Raw("SELECT users.id, users.email, users.fullname, users.image FROM follows LEFT JOIN users ON follows.followers_user_id = users.id WHERE follows.following_user_id = ?", userId).Scan(&usersFollowers).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	return toUserCoreList(usersFollowers), nil

// }