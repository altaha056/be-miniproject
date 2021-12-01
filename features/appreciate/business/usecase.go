package business

import (
	"antonio/features/appreciate"
)

type ratingContentUsecase struct {
	RatingData appreciate.Data
}

func NewRatingBusiness(ratingData appreciate.Data) appreciate.Business {
	return &ratingContentUsecase{RatingData: ratingData}
}

func (alu *ratingContentUsecase) Upvote(articleId, userId int) error {
	err := alu.RatingData.Upvote(articleId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (alu *ratingContentUsecase) Downvote(articleId, userId int) error {
	err := alu.RatingData.Downvote(articleId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (alu *ratingContentUsecase) Rating(userId int) ([]appreciate.ArticleCore, error) {
	articles, err := alu.RatingData.Rating(userId)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (alu *ratingContentUsecase) WhoVote(articleId int) ([]appreciate.UserCore, error) {
	users, err := alu.RatingData.WhoVote(articleId)
	if err != nil {
		return nil, err
	}

	return users, nil
}
