package data

import (
	"antonio/features/appreciate"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type mysqlArticleLikesRepository struct {
	Conn *gorm.DB
}

func NewMysqlArticleLikesRepository(conn *gorm.DB) appreciate.Data {
	return &mysqlArticleLikesRepository{
		Conn: conn,
	}
}

func (alr *mysqlArticleLikesRepository) Upvote(articleId, userId int) error {
	articleLikes := toArticleLikesRecord(appreciate.Core{
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

func (alr *mysqlArticleLikesRepository) Downvote(articleId, userId int) error {

	err := alr.Conn.Where("article_id = ? AND user_id = ?", articleId, userId).Delete(&Rating{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (alr *mysqlArticleLikesRepository) WhoVote(articleId int) ([]appreciate.UserCore, error) {

	var articleLikes []Rating
	err := alr.Conn.Preload(clause.Associations).Joins("User").Where("ratings.article_id = ?", articleId).Find(&articleLikes).Error
	if err != nil {
		return []appreciate.UserCore{}, err
	}
	return toUserCoreList(articleLikes), nil
}

func (alr *mysqlArticleLikesRepository) Rating(userId int) ([]appreciate.ArticleCore, error) {

	var articleLikes []Rating
	err := alr.Conn.Preload(clause.Associations).Joins("Article").Where("ratings.user_id = ?", userId).Find(&articleLikes).Error
	if err != nil {
		return []appreciate.ArticleCore{}, err
	}
	return toArticleCoreList(articleLikes), nil
}
