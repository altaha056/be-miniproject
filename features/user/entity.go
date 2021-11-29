package recruiter

type UserCore struct {
	ID       uint
	Name  string
	Bio      string
	Gender      string
	Email    string
	Password string
	Token    string
}

type Service interface {
	RegisterUser(data UserCore) error
	LoginUser(data UserCore) (UserCore, error)
	GetUsers() ([]UserCore, error)
	GetUserById(data UserCore) (UserCore, error)
}

type Repository interface {
	CreateUser(data UserCore) error
	CheckUser(data UserCore) (UserCore, error)
	GetUsers() ([]UserCore, error)
	GetUserById(data UserCore) (UserCore, error)
}
