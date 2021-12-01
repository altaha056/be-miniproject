package business

import (
	"antonio/features/contents"
)

type contentUsecase struct {
	ArticleData contents.Data
}

func NewArticleBusiness(articleData contents.Data) contents.Business {
	return &contentUsecase{ArticleData: articleData}
}

func (uu *contentUsecase) CreateArticle(data contents.Core, userId int) error {
	tools, err := uu.ArticleData.CreateTags(data.Tools)

	if err != nil {
		return err
	}

	err = uu.ArticleData.CreateArticle(data, userId, tools)
	if err != nil {
		return err
	}
	return nil
}

func (uu *contentUsecase) GetAllArticles() ([]contents.Core, error) {
	articles, err := uu.ArticleData.GetAllArticles()
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (us *contentUsecase) GetArticleById(articleId int) (contents.Core, error) {
	articleData, err := us.ArticleData.GetArticleById(articleId)

	if err != nil {
		return contents.Core{}, err
	}

	return articleData, nil
}

func (uu *contentUsecase) UpdateArticleById(articleId int, data contents.Core, userId int) error {
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

func (uu *contentUsecase) DeleteArticleById(articleId, userId int) error {
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

func (uu *contentUsecase) GetAllUserArticles(userId int) ([]contents.Core, error) {
	articles, err := uu.ArticleData.GetAllUserArticles(userId)
	if err != nil {
		return nil, err
	}

	return articles, nil
}
