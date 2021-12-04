package business

import (
	"antonio/features/appreciate"
	"antonio/features/appreciate/mocks"
	"antonio/features/contents"
	"antonio/features/users"
	"os"
	"testing"
	"time"
)

var (
	likesRepo		mocks.Data
	likesBusiness	appreciate.Business
	usersData		[]users.Core
	articlesData	[]contents.Core
)

func TestMain(m *testing.M) {
	likesBusiness = NewRatingBusiness(&likesRepo)
	usersData = []users.Core{
		{
			ID:       1,
			Email:    "altaha@gmail.com",
			Password: "admin",
		},
		{
			ID:       2,
			Email:    "altaha2@gmail.com",
			Password: "admin",
		},
	}

	articlesData = []contents.Core{
		{
			ID:        3,
			Title:     "monalisa",
			Image:     "monalisa.jpg",
			Content:   "monalisa description",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserId:    3,
			User: contents.UserCore{
				ID:       3,
				Email:    "altaha3@gmail.com",
				Image:    "altaha.jpg",
			},
			Tools: []contents.ToolCore{
				{
					ID:    1,
					Title: "fabel",
				},
				{
					ID:    2,
					Title: "dongeng",
				},
			},
		},
		{
			ID:        4,
			Title:     "pieta",
			Image:     "pieta.jpg",
			Content:   "pieta description",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserId:    3,
			User: contents.UserCore{
				ID:       3,
				Email:    "altaha3@gmail.com",
				Image:    "altaha.jpg",
			},
			Tools: []contents.ToolCore{
				{
					ID:    1,
					Title: "fabel",
				},
				{
					ID:    2,
					Title: "dongeng",
				},
			},
		},
	}

	os.Exit(m.Run())
}