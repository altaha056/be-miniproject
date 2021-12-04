package business

import (
	"antonio/features/comments"
	"antonio/features/comments/mocks"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	commentsRepo		mocks.Data
	commentsBusiness	comments.Business
	commentData			comments.Core
	commentsData		[]comments.Core
)

func TestMain(m *testing.M){
	commentsBusiness=NewCommentsBusiness(&commentsRepo)
	commentData = comments.Core{
		ID: 1,
		Content: "tes komentar 1",
		UserID: 1,
		User: comments.UserCore{
			ID: 1,
			Email: "altaha@alta.com",
			Username: "altaha",
		},
		ArticleID: 1,
		Article: comments.ArticleCore{
			ID: 1,
			Title: "judul konten 1",
			Image: "gambar konten 1",
			Content: "isi deskripsi konten 1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			User: comments.UserCore{
				ID: 1,
				Email: "alta2@alta.com",
				Username: "tes",
			},
			Tags: []comments.TagCore{
				{
					ID: 1,
					Title: "photoshop",
				},
				{
					ID: 2,
					Title: "illustrator",
				},
			},
		},
	}
	
	commentsData = []comments.Core{{ID: 1,
		Content: "tes komentar 3",
		UserID: 1,
		User: comments.UserCore{
			ID: 1,
			Email: "altaha@alta.com",
			Username: "altaha",
		},
		ArticleID: 1,
		Article: comments.ArticleCore{
			ID: 1,
			Title: "judul konten 1",
			Image: "gambar konten 1",
			Content: "isi deskripsi konten 1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			User: comments.UserCore{
				ID: 1,
				Email: "alta2@alta.com",
				Username: "tes",
			},
			Tags: []comments.TagCore{
				{
					ID: 1,
					Title: "photoshop",
				},
				{
					ID: 2,
					Title: "illustrator",
				},
			},
		},},{ID: 1,
		Content: "tes komentar 4",
		UserID: 1,
		User: comments.UserCore{
			ID: 1,
			Email: "altaha@alta.com",
			Username: "altaha",
		},
		ArticleID: 1,
		Article: comments.ArticleCore{
			ID: 1,
			Title: "judul konten 1",
			Image: "gambar konten 1",
			Content: "isi deskripsi konten 1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			User: comments.UserCore{
				ID: 1,
				Email: "alta2@alta.com",
				Username: "tes",
			},
			Tags: []comments.TagCore{
				{
					ID: 1,
					Title: "photoshop",
				},
				{
					ID: 2,
					Title: "illustrator",
				},
			},
		},}}

	os.Exit(m.Run())
}

func TestAddComment(t *testing.T){
	t.Run("comment article success", func(t *testing.T) {
		commentsRepo.On("AddComment", mock.AnythingOfType("content string, articleId int, userId int")).Return(nil).Once()
		err := commentsBusiness.AddComment(commentData.Content, commentData.ArticleID, commentData.UserID)
		assert.Nil(t, err)
	})

	t.Run("comment article fail", func(t *testing.T) {
		commentsRepo.On("AddComment", mock.AnythingOfType("content string, articleId int, userId int")).Return(errors.New("comment article fail")).Once()
		err := commentsBusiness.AddComment(commentData.Content, commentData.ArticleID, commentData.UserID)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "comment article fail")
	})

}

func TestDeleteComment(t *testing.T) {

	t.Run("delete comment success", func(t *testing.T) {
		commentsRepo.On("DeleteComment", mock.AnythingOfType("int")).Return(nil).Once()
		err := commentsBusiness.DeleteComment(1,1)
		assert.Nil(t, err)
	})

	t.Run("delete comment fail", func(t *testing.T) {
		commentsRepo.On("DeleteComment", mock.AnythingOfType("int")).Return(errors.New("delete comment fail")).Once()
		err := commentsBusiness.DeleteComment(1,1)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "delete comment fail")
	})

}

