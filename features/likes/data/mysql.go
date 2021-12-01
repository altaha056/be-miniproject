package data

import (
	"antonio/features/likes"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type mysqlArticleLikesRepository struct {
	Conn *gorm.DB
}

func NewMysqlArticleLikesRepository(conn *gorm.DB) likes.Data {
	return &mysqlArticleLikesRepository{
		Conn: conn,
	}
}

func (alr *mysqlArticleLikesRepository) LikeArticle(articleId, userId int) error {
	articleLikes := toArticleLikesRecord(likes.Core{
		ArticleId: articleId,
		UserId:    userId,
	})
	// err := alr.Conn.Where("article_id = ? AND user_id", articleId, userId).First(&articleLikes).Error
	// if err != nil {
	// 	return errors.New("article is already liked")
	// }
	err := alr.Conn.Create(&articleLikes).Error
	if err != nil {
		return err
	}
	return nil

}

func (alr *mysqlArticleLikesRepository) UnlikeArticle(articleId, userId int) error {

	err := alr.Conn.Where("article_id = ? AND user_id = ?", articleId, userId).Delete(&ArticleLikes{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (alr *mysqlArticleLikesRepository) GetLikingUsers(articleId int) ([]likes.UserCore, error) {

	var articleLikes []ArticleLikes
	err := alr.Conn.Preload(clause.Associations).Joins("User").Where("article_likes.article_id = ?", articleId).Find(&articleLikes).Error
	if err != nil {
		return []likes.UserCore{}, err
	}
	return toUserCoreList(articleLikes), nil
}

func (alr *mysqlArticleLikesRepository) GetLikedArticles(userId int) ([]likes.ArticleCore, error) {

	var articleLikes []ArticleLikes
	err := alr.Conn.Preload(clause.Associations).Joins("Article").Where("article_likes.user_id = ?", userId).Find(&articleLikes).Error
	if err != nil {
		return []likes.ArticleCore{}, err
	}
	return toArticleCoreList(articleLikes), nil
}
