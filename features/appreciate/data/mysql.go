package data

import (
	"antonio/features/appreciate"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type mysqlRatingRepository struct {
	Conn *gorm.DB
}

func NewMysqlRatingRepository(conn *gorm.DB) appreciate.Data {
	return &mysqlRatingRepository{
		Conn: conn,
	}
}

func (alr *mysqlRatingRepository) Upvote(articleId, userId int) error {
	articleRating := toRatingRecord(appreciate.Core{
		ArticleId: articleId,
		UserId:    userId,
	})
	// err := alr.Conn.Where("article_id = ? AND user_id", articleId, userId).First(&articleRating).Error
	// if err != nil {
	// 	return errors.New("article is already liked")
	// }
	err := alr.Conn.Create(&articleRating).Error
	if err != nil {
		return err
	}
	return nil

}

func (alr *mysqlRatingRepository) Downvote(articleId, userId int) error {

	err := alr.Conn.Where("article_id = ? AND user_id = ?", articleId, userId).Delete(&Rating{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (alr *mysqlRatingRepository) WhoVote(articleId int) ([]appreciate.UserCore, error) {

	var articleRating []Rating
	err := alr.Conn.Preload(clause.Associations).Joins("User").Where("ratings.article_id = ?", articleId).Find(&articleRating).Error
	if err != nil {
		return []appreciate.UserCore{}, err
	}
	return toUserCoreList(articleRating), nil
}

func (alr *mysqlRatingRepository) Rating(userId int) ([]appreciate.ArticleCore, error) {

	var articleRating []Rating
	err := alr.Conn.Preload(clause.Associations).Joins("Article").Where("ratings.user_id = ?", userId).Find(&articleRating).Error
	if err != nil {
		return []appreciate.ArticleCore{}, err
	}
	return toArticleCoreList(articleRating), nil
}
