package business

import (
	"antonio/features/contents"
	"antonio/features/contents/mocks"
	"os"
	"testing"
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