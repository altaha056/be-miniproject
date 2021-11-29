package factory

import (
	"antonio/driver"
	userData "antonio/features/user/data"
	userPresenter "antonio/features/user/presentation"
	userService "antonio/features/user/service"
)

type factoryPresenter struct{
	UserPresentation	userPresenter.UserHandler
}

func Init() factoryPresenter{
	userData:=userData.NewMysqlUserRepository(driver.DB)
	userService:=userService.NewUserService(userData)

	return factoryPresenter{
		UserPresentation:  *userPresenter.NewUserHandler(userService),
	}
}