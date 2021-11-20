package posts

import "time"

type Core struct {
	ID    int
	UrlImg    string
	Caption   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct{
	Username	string
	Email		string
}

type Bussiness interface {
	CreateData(data Core) (resp Core, err error)
	GetAllData(search string) (resp []Core)
	// continue the function abtraction
}

type Data interface {
	InsertData(data Core) (resp Core, err error)
	SelectData(title string) (resp []Core)
}


