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

func toUserRecord(user contents.UserCore) User {
	return User{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}
}

func toTagRecord(tag contents.TagCore) Tag {
	return Tag{
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
		Tags:      toTagsRecordList(article.Tags),
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
		Tags:      toTagsCoreList(article.Tags),
	}
}

func toUserCore(user User) contents.UserCore {
	return contents.UserCore{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}
}

func toTagCore(tag Tag) contents.TagCore {
	return contents.TagCore{
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

func toTagsCoreList(tList []Tag) []contents.TagCore {
	convertedTag := []contents.TagCore{}

	for _, tag := range tList {
		convertedTag = append(convertedTag, toTagCore(tag))
	}

	return convertedTag
}

func toTagsRecordList(tList []contents.TagCore) []Tag {
	convertedUser := []Tag{}

	for _, tag := range tList {
		convertedUser = append(convertedUser, toTagRecord(tag))
	}

	return convertedUser
}
