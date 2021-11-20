package data

import (
	"mpbe/features/posts"

	"gorm.io/gorm"
)

type mysqlPostRepository struct {
	Conn *gorm.DB
}
func NewPostRepository(conn *gorm.DB)posts.Data{
	return &mysqlPostRepository{
		Conn: conn,
	}
}

func (po *mysqlPostRepository)InsertData(data posts.Core)(resp posts.Core, err error){
	return posts.Core{}, nil
}

func (po *mysqlPostRepository)SelectData(caption string)(resp[]posts.Core){
	var record []Post
	if err:=po.Conn.Preload("TagPost").Find(&record).Error; err!=nil{
		return []posts.Core{}
	}
	return toCoreList(record)
}