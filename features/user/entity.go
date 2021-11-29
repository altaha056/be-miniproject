package user

type UserCore struct{
	Id			uint
	Name		string
	Bio			string
	Gender		string
	Email		string
	Password	string
	Token		string
}

type Service interface{
	RegisterUser(data UserCore) (err error)
	LoginUser(data UserCore) (user UserCore, err error)
	GetUsers(data UserCore) (user []UserCore, err error)
	GetUserByName(name string) (user UserCore, err error)
}

type Repository interface{
	InsertData(data UserCore) (err error)
	CheckUser(data UserCore) (user UserCore, err error)
	GetData(UserCore) (user []UserCore,err error)
	GetDataByName(name string) (user UserCore,err error)

}