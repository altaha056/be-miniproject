package business

import (
	"antonio/features/authentication"
	"antonio/features/users"
	"antonio/middlewares"
)

type authUsecase struct {
	AuthData auth.Data
}

func NewAuthBusiness(authData auth.Data) auth.Business {
	return &authUsecase{AuthData: authData}
}

func (au *authUsecase) Login(data users.Core) (accessTokenCore auth.Core, /*refreshTokenCore auth.Core,*/ err error) {
	user, err := au.AuthData.TokenAuthentication(data)
	if err != nil {
		return auth.Core{}, /*auth.Core{},*/ err
	}
	accessToken, err := middlewares.CreateToken(user.ID)
	if err != nil {
		return auth.Core{}, err
	}
	accessTokenCore = auth.Core{
		Token: accessToken,
	}

	return accessTokenCore, nil
}


