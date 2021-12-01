package presentation

import (
	"net/http"
	"antonio/features/appreciate"
	"antonio/features/appreciate/presentation/response"
	"antonio/middlewares"

	"strconv"

	"github.com/labstack/echo/v4"
)

type RatingHandler struct {
	RatingBusiness appreciate.Business
}

func NewRatingHandler(ratingBusiness appreciate.Business) *RatingHandler {
	return &RatingHandler{RatingBusiness: ratingBusiness}
}

func (uh *RatingHandler) LikeArticle(e echo.Context) error {

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
	err = uh.RatingBusiness.Upvote(articleId, userId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  "upvoted!",
	})
}

func (uh *RatingHandler) UnlikeArticle(e echo.Context) error {
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
	err = uh.RatingBusiness.Downvote(articleId, userId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "downvoted!",
	})

}

func (alh *RatingHandler) GetLikedArticles(e echo.Context) error {
	userId, err := strconv.Atoi(e.Param("userId"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err":     err.Error(),
		})
	}
	data, err := alh.RatingBusiness.Rating(userId)
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

func (alh *RatingHandler) GetLikingUsers(e echo.Context) error {
	articleId, err := strconv.Atoi(e.Param("articleId"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err":     err.Error(),
		})
	}
	data, err := alh.RatingBusiness.WhoVote(articleId)
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
