package posts

import "time"

type Core struct {
	IdPost    int
	UrlImg    string
	Caption   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct{
	Username	string
	Email		string
}

