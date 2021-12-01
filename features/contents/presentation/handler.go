package presentation

import (
	"net/http"
	"antonio/features/contents"
	"antonio/features/contents/presentation/request"
	"antonio/features/contents/presentation/response"
	"antonio/middlewares"

	"strconv"

	"github.com/labstack/echo/v4"
)

type ArticleHandler struct {
	ArticleBusiness contents.Business
}

func NewArticleHandler(articleBusiness contents.Business) *ArticleHandler {
	return &ArticleHandler{ArticleBusiness: articleBusiness}
}

func (uh *ArticleHandler) CreateArticleHandler(e echo.Context) error {
	articleData := request.ArticleRequest{}

	err := e.Bind(&articleData)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"err":     err.Error(),
		})
	}

	userId, err := middlewares.VerifyAccessToken(e)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err":     err.Error(),
		})
	}
	err = uh.ArticleBusiness.CreateArticle(articleData.ToArticleCore(), userId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "content uploaded!",
	})
}

func (ah *ArticleHandler) GetAllArticleHandler(e echo.Context) error {
	data, err := ah.ArticleBusiness.GetAllArticles()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   response.ToArticleResponseList(data),
	})

}

func (ah *ArticleHandler) GetArticleByIdHandler(e echo.Context) error {
	contentId, err := strconv.Atoi(e.Param("contentId"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"err":     err.Error(),
		})
	}

	data, err := ah.ArticleBusiness.GetArticleById(contentId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   response.ToArticleResponse(data),
	})

}

func (ah *ArticleHandler) UpdateArticleByIdHandler(e echo.Context) error {
	contentId, _ := strconv.Atoi(e.Param("contentId"))
	articleData := request.ArticleRequest{}
	err := e.Bind(&articleData)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"err":     err.Error(),
		})
	}

	userId, err := middlewares.VerifyAccessToken(e)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err":     err.Error(),
		})
	}

	err = ah.ArticleBusiness.UpdateArticleById(contentId, articleData.ToArticleCore(), userId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "content updated!",
	})

}

func (uh *ArticleHandler) DeleteArticleByIdHandler(e echo.Context) error {
	contentId, _ := strconv.Atoi(e.Param("contentId"))

	userId, err := middlewares.VerifyAccessToken(e)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err":     err.Error(),
		})
	}

	err = uh.ArticleBusiness.DeleteArticleById(contentId, userId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "content deleted!",
	})

}

func (ah *ArticleHandler) GetAllUserArticlesHandler(e echo.Context) error {

	userId, err := strconv.Atoi(e.Param("userId"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err":     err.Error(),
		})
	}
	data, err := ah.ArticleBusiness.GetAllUserArticles(userId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   response.ToArticleResponseList(data),
	})

}
