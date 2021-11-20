package data

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	IdPost uint
	UrlImg string `gorm:"column:url_img"`
	Caption string `gorm:"column:caption"`
	IdOwner uint `gorm:"column:id_owner"`
	CategoryID uint
	TagPost []TagPost    `gorm:"many2many:posts_tag"`
	
}

type TagPost struct{
	gorm.Model	
	Name   string `gorm:"type:varchar(50)"`
}

///DTO