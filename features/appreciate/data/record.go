package data

import (
	"antonio/features/appreciate"
	"time"

	"gorm.io/gorm"
)

type Rating struct {
	gorm.Model
	ArticleID int
	Article   Content
	UserID    int
	User      User
}

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
	Tools     []Tool `gorm:"many2many:content_tool;"`
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

func toRatingRecord(data appreciate.Core) Rating {
	return Rating{
		UserID:    data.UserId,
		ArticleID: data.ArticleId,
	}
}

func toUserRecord(user appreciate.UserCore) User {
	return User{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}
}

func toTagRecord(tool appreciate.ToolCore) Tool {
	return Tool{
		ID:    tool.ID,
		Title: tool.Title,
	}
}

func toArticleRecord(article appreciate.ArticleCore) Content {
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

func toArticleCore(article Content) appreciate.ArticleCore {
	return appreciate.ArticleCore{
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

func toUserCore(user User) appreciate.UserCore {
	return appreciate.UserCore{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}
}

func toTagCore(tool Tool) appreciate.ToolCore {
	return appreciate.ToolCore{
		ID:    tool.ID,
		Title: tool.Title,
	}
}

func toUserCoreList(aList []Rating) []appreciate.UserCore {
	convertedUser := []appreciate.UserCore{}

	for _, user := range aList {
		convertedUser = append(convertedUser, toUserCore(user.User))
	}

	return convertedUser
}
func toArticleCoreList(aList []Rating) []appreciate.ArticleCore {
	convertedArticle := []appreciate.ArticleCore{}

	for _, article := range aList {
		convertedArticle = append(convertedArticle, toArticleCore(article.Article))
	}

	return convertedArticle
}

func toTagsCoreList(tList []Tool) []appreciate.ToolCore {
	convertedTag := []appreciate.ToolCore{}

	for _, tool := range tList {
		convertedTag = append(convertedTag, toTagCore(tool))
	}

	return convertedTag
}

func toTagsRecordList(tList []appreciate.ToolCore) []Tool {
	convertedUser := []Tool{}

	for _, tool := range tList {
		convertedUser = append(convertedUser, toTagRecord(tool))
	}

	return convertedUser
}
