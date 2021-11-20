package data

import (
	"fmt"
	"mpbe/features/posts"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	UrlImg string `gorm:"column:url_img"`
	Caption string `gorm:"column:caption"`
	IdOwner uint
	TagPost []TagPost    `gorm:"many2many:posts_tag"`
	
}



type TagPost struct{
	gorm.Model	
	Name   string `gorm:"type:varchar(50)"`
}

///DTO

func (a *Post) toCore() posts.Core{
	fmt.Printf("%+v\n\n", a)
	return posts.Core{
		ID: int(a.ID),
		UrlImg: a.UrlImg,
		Caption: a.Caption,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

func toCoreList(resp []Post) []posts.Core{
	a:=[]posts.Core{}
	for key:=range resp{
		a=append(a, resp[key].toCore())
	}
	return a
}