package data

import (
	"antonio/features/articles"
	"antonio/features/appreciate"
	"time"

	"gorm.io/gorm"
)

type Rating struct {
	gorm.Model
	ArticleID int
	Article   Article
	UserID    int
	User      User
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

func toArticleLikesRecord(data appreciate.Core) Rating {
	return Rating{
		UserID:    data.UserId,
		ArticleID: data.ArticleId,
	}
}

func toUserRecord(user articles.UserCore) User {
	return User{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}
}

func toTagRecord(tag articles.TagCore) Tag {
	return Tag{
		ID:    tag.ID,
		Title: tag.Title,
	}
}

func toArticleRecord(article articles.Core) Article {
	return Article{
		ID:        article.ID,
		Title:     article.Title,
		Image:     article.Image,
		Content:   article.Content,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
		UserID:    article.UserId,
		User:      toUserRecord(article.User),
		Tags:      toTagsRecordList(article.Tags),
	}
}

func toArticleCore(article Article) appreciate.ArticleCore {
	return appreciate.ArticleCore{
		ID:        article.ID,
		Title:     article.Title,
		Image:     article.Image,
		Content:   article.Content,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
		User:      toUserCore(article.User),
		Tags:      toTagsCoreList(article.Tags),
	}
}

func toUserCore(user User) appreciate.UserCore {
	return appreciate.UserCore{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}
}

func toTagCore(tag Tag) appreciate.TagCore {
	return appreciate.TagCore{
		ID:    tag.ID,
		Title: tag.Title,
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

func toTagsCoreList(tList []Tag) []appreciate.TagCore {
	convertedTag := []appreciate.TagCore{}

	for _, tag := range tList {
		convertedTag = append(convertedTag, toTagCore(tag))
	}

	return convertedTag
}

func toTagsRecordList(tList []articles.TagCore) []Tag {
	convertedUser := []Tag{}

	for _, tag := range tList {
		convertedUser = append(convertedUser, toTagRecord(tag))
	}

	return convertedUser
}
