package factory

import (
	"antonio/driver"
	userData "antonio/features/user/data"
	userPresentaion "antonio/features/user/presentation"
	userService "antonio/features/user/service"
)

type factoryPresenter struct{
	UserPresentation userPresentaion.UserHandler
}

func Init() factoryPresenter{
	userData := userData.NewMysqlUserRepository(driver.DB)
	userService := userService.NewUserService(userData)

	return factoryPresenter{
		UserPresentation: *userPresentaion.NewUserHandler(userService),
	}
}

