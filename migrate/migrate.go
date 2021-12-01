package migrate

import (
	"antonio/db"
	articles "antonio/features/articles/data"
	auth "antonio/features/authentication/data"
	comments "antonio/features/comments/data"
	likes "antonio/features/likes/data"
	users "antonio/features/users/data"
)

func AutoMigrate() {
	db.DB.AutoMigrate(
		&users.User{},
		&auth.Authentication{},
		&articles.Article{},
		&articles.Tag{},
		&likes.ArticleLikes{},
		&comments.Comment{},
	)
}
