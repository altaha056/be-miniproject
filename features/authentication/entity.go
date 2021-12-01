package auth

import "antonio/features/users"

type Core struct {
	Token string
}

type Business interface {
	Login(data users.Core) (accessTokenCore Core, err error)
}

type Data interface {
	TokenAuthentication(data users.Core) (users.Core, error)
}
