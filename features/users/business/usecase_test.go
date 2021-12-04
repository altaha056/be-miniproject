package business

import (
	"antonio/features/users"
	"antonio/features/users/mocks"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userRepo 	mocks.Data
	userServ	users.Business
	usersData	[]users.Core
	userData	users.Core
	userLogin	users.Core
)

func TestMain(m *testing.M){
	userServ = NewUserBusiness(&userRepo)
	usersData = []users.Core{
		{
			ID: 1,
			Username: "Altaha",
			Email: "altaha@gmail.com",
			Password: "altahapass777",
			Bio: "another bio",
			Gender: "male",
		},
	}

	userData = users.Core{
		
			ID: 1,
			Username: "Altaha",
			Email: "altaha@gmail.com",
			Password: "altahapass777",
			Bio: "another bio",
			Gender: "male",
		
	}

	userLogin = users.Core{
		Email: "altaha@gmail.com",
		Password: "123123",
	}

	os.Exit(m.Run())
}

func TestGetUsers(t *testing.T){
	t.Run("validate get users", func(t *testing.T) {
		userRepo.On("GetAllUsers", mock.AnythingOfType("users.Core")).Return(usersData, nil).Once()
		resp, err := userServ.GetAllUsers()
		assert.Nil(t, err)
		assert.Equal(t, len(resp), 1)
	})

	t.Run("error get users", func(t *testing.T) {
		userRepo.On("GetAllUsers", mock.AnythingOfType("users.Core")).Return(nil, errors.New("error on db"))
		resp, err := userServ.GetAllUsers()
		assert.NotNil(t, err)
		assert.Nil(t, resp)
	})
}

func TestRegisterUser(t *testing.T){
	t.Run("register user success", func(t *testing.T){
		userRepo.On("CreateUser", mock.AnythingOfType("string")).Return(false, nil).Once()
	})
}