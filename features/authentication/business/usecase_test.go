package business

import (
	auth "antonio/features/authentication"
	"antonio/features/authentication/mocks"
	"antonio/features/users"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	authRepo mocks.Data
	authBusiness auth.Business
	authData auth.Core
	userData users.Core
)

func TestMain(m *testing.M){
	authBusiness = NewAuthBusiness(&authRepo)
	authData=auth.Core{
		Token: "123123123",
	}
	userData=users.Core{
		ID: 1,
		Username: "altaha",
		Email: "altaha@pb.com",
		Password: "asdasdasd",
		Bio: "test bio",
		Gender: "male",
	}
	os.Exit(m.Run())
}

func TestLogin(t *testing.T){
	t.Run("verify success", func(t *testing.T) {
		authRepo.On("TokenAuthentication", mock.AnythingOfType("users.Core")).Return(1, nil).Once()
		userId, err:= authBusiness.Login(userData)
		assert.Nil(t, err)
		assert.Equal(t, userId, 1)
	})
}