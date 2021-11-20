package migrate

import (
	"mpbe/config"
	m_posts "mpbe/features/posts/data"
)


func AutoMigrate() {
	config.DB.AutoMigrate(
		&m_posts.Post{},
	)
}