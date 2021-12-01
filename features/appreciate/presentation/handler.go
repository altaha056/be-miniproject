package presentation

import (
	"net/http"
	"antonio/features/appreciate"
	"antonio/features/appreciate/presentation/response"
	"antonio/middlewares"

	"strconv"

	"github.com/labstack/echo/v4"
)

type ArticleLikesHandler struct {
	ArticleLikesBusiness appreciate.Business
}

func NewArticleLikesHandler(articleLikesBusiness appreciate.Business) *ArticleLikesHandler {
	return &ArticleLikesHandler{ArticleLikesBusiness: articleLikesBusiness}
}

func (uh *ArticleLikesHandler) LikeArticle(e echo.Context) error {

	articleId, err := strconv.Atoi(e.Param("articleId"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err":     err.Error(),
		})
	}
	userId, err := middlewares.VerifyAccessToken(e)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err":     err.Error(),
		})
	}
	err = uh.ArticleLikesBusiness.Upvote(articleId, userId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  "upvoted!",
	})
}

func (uh *ArticleLikesHandler) UnlikeArticle(e echo.Context) error {
	articleId, err := strconv.Atoi(e.Param("articleId"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err":     err.Error(),
		})
	}
	userId, err := middlewares.VerifyAccessToken(e)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err":     err.Error(),
		})
	}
	err = uh.ArticleLikesBusiness.Downvote(articleId, userId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "downvoted!",
	})

}

func (alh *ArticleLikesHandler) GetLikedArticles(e echo.Context) error {
	userId, err := strconv.Atoi(e.Param("userId"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err":     err.Error(),
		})
	}
	data, err := alh.ArticleLikesBusiness.Rating(userId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   response.ToArticleResponseList(data),
	})

}

func (alh *ArticleLikesHandler) GetLikingUsers(e echo.Context) error {
	articleId, err := strconv.Atoi(e.Param("articleId"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err":     err.Error(),
		})
	}
	data, err := alh.ArticleLikesBusiness.WhoVote(articleId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   response.ToUserResponseList(data),
	})

}
