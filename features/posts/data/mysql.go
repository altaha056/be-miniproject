package data

import "gorm.io/gorm"

type mysqlPostRepository struct {
	Conn *gorm.DB
}
func NewPostRepository(conn *gorm.DB){
	
}