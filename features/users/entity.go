package users

type Core struct {
	ID       	int
	Username	string
	Email    	string
	Password 	string
	Bio			string
	Gender		string
}

type Business interface {
	RegisterUser(data Core) (userId int, err error)
	GetAllUsers() ([]Core, error)
	GetUserById(id int) (Core, error)
}

type Data interface {
	CreateUser(data Core) (userId int, err error)
	GetAllUsers() ([]Core, error)
	GetUserById(userId int) (Core, error)
}
