package data

import (
	"fmt"
	"antonio/features/contents"

	"gorm.io/gorm"
)

type mysqlArticleRepository struct {
	Conn *gorm.DB
}

func NewMysqlArticleRepository(conn *gorm.DB) contents.Data {
	return &mysqlArticleRepository{
		Conn: conn,
	}
}

func (ar *mysqlArticleRepository) CreateTags(tags []contents.ToolCore) ([]contents.ToolCore, error) {
	tagsTitle := make([]string, 0, len(tags))
	articleTags := []Tool{}
	for _, tag := range tags {
		err := ar.Conn.Where("title = ?", tag.Title).First(&Tool{}).Error
		if err != nil {
			err := ar.Conn.Create(&Tool{Title: tag.Title}).Error
			if err != nil {
				return []contents.ToolCore{}, err
			}
		}
		tagsTitle = append(tagsTitle, tag.Title)
	}
	err := ar.Conn.Where("title IN ?", tagsTitle).Find(&articleTags).Error
	if err != nil {
		return []contents.ToolCore{}, err
	}
	return toTagsCoreList(articleTags), nil

}

func (ar *mysqlArticleRepository) CreateArticle(data contents.Core, userId int, tools []contents.ToolCore) error {

	data.UserId = userId
	data.Tools = tools

	recordData := toArticleRecord(data)
	err := ar.Conn.Create(&recordData).Error
	if err != nil {
		return err
	}
	return nil
}

func (ar *mysqlArticleRepository) GetAllArticles() ([]contents.Core, error) {

	articles := []Content{}
	err := ar.Conn.Joins("User").Preload("Tools").Find(&articles).Error
	if err != nil {
		return toArticleCoreList([]Content{}), err
	}
	return toArticleCoreList(articles), nil
}

func (ar *mysqlArticleRepository) GetArticleById(articleId int) (contents.Core, error) {

	article := Content{}
	err := ar.Conn.Joins("User").Preload("Tools").First(&article, articleId).Error
	if err != nil {
		return toArticleCore(Content{}), err
	}
	return toArticleCore(article), nil
}

func (ar *mysqlArticleRepository) UpdateArticleById(articleId int, data contents.Core) error {

	article := toArticleRecord(data)
	article.ID = articleId
	err := ar.Conn.Save(&article).Error
	if err != nil {
		return err
	}
	return nil
}

func (ar *mysqlArticleRepository) DeleteArticleById(articleId int) error {

	err := ar.Conn.Delete(&Content{}, articleId).Error
	if err != nil {
		return err
	}
	return nil
}

func (ar *mysqlArticleRepository) VerifyArticleOwner(articleId int, userId int) error {

	fmt.Print(articleId, userId)
	err := ar.Conn.Where("id = ? AND user_id = ?", articleId, userId).First(&Content{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (ar *mysqlArticleRepository) GetAllUserArticles(userId int) ([]contents.Core, error) {

	articles := []Content{}
	err := ar.Conn.Joins("User").Preload("Tools").Where("user_id = ?", userId).Find(&articles).Error
	if err != nil {
		return toArticleCoreList([]Content{}), err
	}
	return toArticleCoreList(articles), nil
}
