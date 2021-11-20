package business

import (
	"mpbe/features/posts"

	"github.com/go-playground/validator/v10"
)

type postsUsecase struct {
	postData posts.Data
	validate *validator.Validate
}

func NewPostBusiness(pData posts.Data)posts.Bussiness{
	return &postsUsecase{
		postData: pData,
		validate: validator.New(),
	}
}

func (pu *postsUsecase) CreateData(data posts.Core)(resp posts.Core, err error){
	return posts.Core{},nil
}

func (pu *postsUsecase) GetAllData(search string)(resp []posts.Core){
	resp=pu.postData.SelectData(search)
	return
}