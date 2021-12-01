package business

import (
	"antonio/features/contents"
)

type articleUsecase struct {
	ArticleData contents.Data
}

func NewArticleBusiness(articleData contents.Data) contents.Business {
	return &articleUsecase{ArticleData: articleData}
}

func (uu *articleUsecase) CreateArticle(data contents.Core, userId int) error {
	tags, err := uu.ArticleData.CreateTags(data.Tags)

	if err != nil {
		return err
	}

	err = uu.ArticleData.CreateArticle(data, userId, tags)
	if err != nil {
		return err
	}
	return nil
}

func (uu *articleUsecase) GetAllArticles() ([]contents.Core, error) {
	articles, err := uu.ArticleData.GetAllArticles()
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (us *articleUsecase) GetArticleById(articleId int) (contents.Core, error) {
	articleData, err := us.ArticleData.GetArticleById(articleId)

	if err != nil {
		return contents.Core{}, err
	}

	return articleData, nil
}

func (uu *articleUsecase) UpdateArticleById(articleId int, data contents.Core, userId int) error {
	err := uu.ArticleData.VerifyArticleOwner(articleId, userId)
	if err != nil {
		return err
	}
	err = uu.ArticleData.UpdateArticleById(articleId, data)
	if err != nil {
		return err
	}

	return nil
}

func (uu *articleUsecase) DeleteArticleById(articleId, userId int) error {
	err := uu.ArticleData.VerifyArticleOwner(articleId, userId)
	if err != nil {
		return err
	}

	err = uu.ArticleData.DeleteArticleById(articleId)
	if err != nil {
		return err
	}

	return nil
}

func (uu *articleUsecase) GetAllUserArticles(userId int) ([]contents.Core, error) {
	articles, err := uu.ArticleData.GetAllUserArticles(userId)
	if err != nil {
		return nil, err
	}

	return articles, nil
}
