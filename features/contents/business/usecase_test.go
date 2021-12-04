package business

import (
	"antonio/features/contents"
	"antonio/features/contents/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	cService 	contents.Business
	cRepository	mocks.Data
	cData		contents.Core
	csData		[]contents.Core
)

func TestMain(m *testing.M){
	cService = NewArticleBusiness(&cRepository)

	cData = contents.Core{
		Title: "Last Supper",
		Content: "A masterpiece from Leonardo Da Vinci",
		Tools: []contents.ToolCore{
			{Title: "photoshop"},
			{Title: "Lightroom"},
		},
	}

	csData=[]contents.Core{
		{
			ID:			1,
			Title:		"Fresko in Sistina Chapel",
			Content:	"i don't know how to describe how amazing is this masterpiece",
			Tools:		[]contents.ToolCore{
				{Title: "adobe illustrator"},
				{Title: "premiere pro"},
			},
		},
	}

	os.Exit(m.Run())
}

func TestCreateContent(t *testing.T){
	t.Run("create content success", func(t *testing.T) {
		cRepository.On("CreateArticle", mock.AnythingOfType("contents.Core")).Return(nil).Once()
		err:=cService.CreateArticle(cData, cData.UserId)
		assert.Nil(t, err)
	})
}

func TestGetContent(t *testing.T){
	t.Run("get content success", func(t *testing.T) {
		cRepository.On("GetAllArticles", mock.AnythingOfType("contents.Core")).Return(csData, nil).Once()
	})
}

func TestDeleteContent(t *testing.T){
	t.Run("delete content", func(t *testing.T) {
		cRepository.On("GetArticleById", mock.AnythingOfType("int")).Return(cData,nil).Once()
		cRepository.On("DeleteArticleById",mock.AnythingOfType("contents.Core")).Return(nil).Once()
		err := cService.DeleteArticleById(cData.ID, cData.UserId)
		assert.Nil(t, err)
	})
}