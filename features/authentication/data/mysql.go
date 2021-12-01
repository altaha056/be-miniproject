package data

import (
	"errors"
	"antonio/features/authentication"
	"antonio/features/users"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type mysqlAuthRepository struct {
	Conn *gorm.DB
}

func NewMysqlAuthRepository(conn *gorm.DB) auth.Data {
	return &mysqlAuthRepository{
		Conn: conn,
	}
}

func (ur *mysqlAuthRepository) TokenAuthentication(data users.Core) (users.Core, error) {

	recordData := toUserRecord(data)
	ur.Conn.Where("email = ?", data.Email).First(&recordData)
	if recordData.Password == "" && recordData.ID == 0 {
		return users.Core{
			ID: 0,
		}, errors.New("incorrect email or password")
	}
	if bcrypt.CompareHashAndPassword([]byte(recordData.Password), []byte(data.Password)) != nil {
		return users.Core{
			ID: 0,
		}, errors.New("incorrect email or password")
	}

	return toUserCore(recordData), nil
}
