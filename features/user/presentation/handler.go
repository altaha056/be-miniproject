package presentation

import (
	"antonio/features/user"
	"antonio/features/user/presentation/request"
	"antonio/features/user/presentation/response"
	"net/http"
	"strconv"
	"github.com/labstack/echo"
)

type UserHandler struct {
	userService user.Service
}

func NewUserHandler(uService user.Service) *UserHandler{
	return &UserHandler{uService}
}

func (uhn *UserHandler)RegisterUserHandler(e echo.Context)error{
	userData:=request.UserRequest{}
	err:=e.Bind(&userData)
	if err!=nil{
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":err.Error(),
		})
	}
	err=uhn.userService.RegisterUser(userData.TouserCore())
	if err!=nil{
		return e.JSON(http.StatusInternalServerError,map[string]interface{}{
			"message":err.Error(),
		})
	}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message":"success",
		})
}


func (uh *UserHandler) GetUsersHandler(e echo.Context) error {
	
	data, err := uh.userService.GetUsers(user.UserCore{})
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToUserResponseList(data),
	})

}

func (uh *UserHandler) LoginUserHandler(e echo.Context) error {
	userAuth := request.UserAuth{}
	e.Bind(&userAuth)
	data, err := uh.userService.LoginUser(userAuth.TouserCore())

	if err != nil {
		return e.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToUserLoginResponse(data),
	})

}

func (uh *UserHandler) GetUserByIdHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	data, err := uh.userService.GetUserById(id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToUserResponse(data),
	})

}
