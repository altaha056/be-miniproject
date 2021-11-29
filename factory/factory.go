package factory

import (
	"antonio/driver"
	userData "antonio/features/user/data"
	userPresentaion "antonio/features/user/presentation"
	userBusiness "antonio/features/user/business"
)

type factoryPresenter struct{
	UserPresentation userPresentaion.UserHandler
}

func Init() factoryPresenter{
	userData := userData.NewMysqlUserRepository(driver.DB)
	userBusiness := userBusiness.NewUserService(userData)

	return factoryPresenter{
		UserPresentation: *userPresentaion.NewUserHandler(userBusiness),
	}
}

