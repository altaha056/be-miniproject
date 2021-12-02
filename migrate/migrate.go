package migrate

import (
	"antonio/db"
	contents "antonio/features/contents/data"
	auth "antonio/features/authentication/data"
	comments "antonio/features/comments/data"
	appreciate "antonio/features/appreciate/data"
	users "antonio/features/users/data"
)

func AutoMigrate() {
	db.DB.AutoMigrate(
		&users.User{},
		&auth.Authentication{},
		&contents.Content{},
		&contents.Tool{},
		&appreciate.Rating{},
		&comments.Comment{},
	)
}
