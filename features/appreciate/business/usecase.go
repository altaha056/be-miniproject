package business

import (
	"antonio/features/appreciate"
)

type articleLikesUsecase struct {
	LikeData appreciate.Data
}

func NewArticleLikesBusiness(articleLikesData appreciate.Data) appreciate.Business {
	return &articleLikesUsecase{LikeData: articleLikesData}
}

func (alu *articleLikesUsecase) Upvote(articleId, userId int) error {
	err := alu.LikeData.Upvote(articleId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (alu *articleLikesUsecase) Downvote(articleId, userId int) error {
	err := alu.LikeData.Downvote(articleId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (alu *articleLikesUsecase) Rating(userId int) ([]appreciate.ArticleCore, error) {
	articles, err := alu.LikeData.Rating(userId)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (alu *articleLikesUsecase) WhoVote(articleId int) ([]appreciate.UserCore, error) {
	users, err := alu.LikeData.WhoVote(articleId)
	if err != nil {
		return nil, err
	}

	return users, nil
}
