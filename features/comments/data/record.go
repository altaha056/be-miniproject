package data

import (
	"antonio/features/articles"
	"antonio/features/comments"
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID        int
	Content   string
	UserID    int
	User      User
	ArticleID int
	Article   Article
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Article struct {
	gorm.Model
	ID        int
	Title     string
	Image     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int
	User      User
	Tags      []Tag `gorm:"many2many:article_tags;"`
}

type User struct {
	gorm.Model
	ID       int
	Email    string
	Username string
}

type Tag struct {
	gorm.Model
	ID    int
	Title string `gorm:"unique"`
}

func toCommentsRecord(data comments.Core) Comment {
	return Comment{
		Content:   data.Content,
		UserID:    data.UserID,
		ArticleID: data.ArticleID,
	}
}

func toUserRecord(user articles.UserCore) User {
	return User{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}
}

func toUserCore(user User) comments.UserCore {
	return comments.UserCore{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}
}

func toCommentCore(comment Comment) comments.Core {
	return comments.Core{
		ID:        comment.ID,
		Content:   comment.Content,
		User:      toUserCore(comment.User),
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}
}

func toCommentCoreList(cList []Comment) []comments.Core {
	convertedComment := []comments.Core{}

	for _, comment := range cList {
		convertedComment = append(convertedComment, toCommentCore(comment))
	}

	return convertedComment
}
