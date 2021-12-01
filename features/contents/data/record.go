package data

import (
	"antonio/features/contents"
	"time"

	"gorm.io/gorm"
)

type Content struct {
	gorm.Model
	ID        int
	Title     string
	Image     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int
	User      User
	Tools     []Tool `gorm:"many2many:content_tools;"`
}

type User struct {
	gorm.Model
	ID       int
	Email    string
	Username string
}

type Tool struct {
	gorm.Model
	ID    int
	Title string `gorm:"unique"`
}

func toUserRecord(user contents.UserCore) User {
	return User{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}
}

func toTagRecord(tag contents.ToolCore) Tool {
	return Tool{
		ID:    tag.ID,
		Title: tag.Title,
	}
}

func toArticleRecord(article contents.Core) Content {
	return Content{
		ID:        article.ID,
		Title:     article.Title,
		Image:     article.Image,
		Content:   article.Content,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
		UserID:    article.UserId,
		User:      toUserRecord(article.User),
		Tools:     toTagsRecordList(article.Tools),
	}
}

func toArticleCore(article Content) contents.Core {
	return contents.Core{
		ID:        article.ID,
		Title:     article.Title,
		Image:     article.Image,
		Content:   article.Content,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
		User:      toUserCore(article.User),
		Tools:     toTagsCoreList(article.Tools),
	}
}

func toUserCore(user User) contents.UserCore {
	return contents.UserCore{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}
}

func toTagCore(tag Tool) contents.ToolCore {
	return contents.ToolCore{
		ID:    tag.ID,
		Title: tag.Title,
	}
}

func toArticleCoreList(aList []Content) []contents.Core {
	convertedArticle := []contents.Core{}

	for _, article := range aList {
		convertedArticle = append(convertedArticle, toArticleCore(article))
	}

	return convertedArticle
}

func toTagsCoreList(tList []Tool) []contents.ToolCore {
	convertedTag := []contents.ToolCore{}

	for _, tag := range tList {
		convertedTag = append(convertedTag, toTagCore(tag))
	}

	return convertedTag
}

func toTagsRecordList(tList []contents.ToolCore) []Tool {
	convertedUser := []Tool{}

	for _, tag := range tList {
		convertedUser = append(convertedUser, toTagRecord(tag))
	}

	return convertedUser
}
