package presentation

import (
	"net/http"
	"antonio/features/comments"
	"antonio/features/comments/presentation/request"
	"antonio/features/comments/presentation/response"
	"antonio/middlewares"
	"strconv"
	"github.com/labstack/echo/v4"
)

type CommentHandler struct {
	CommentBusiness comments.Business
}

func NewCommentHandler(commentBusiness comments.Business) *CommentHandler {
	return &CommentHandler{CommentBusiness: commentBusiness}
}

func (uh *CommentHandler) AddComment(e echo.Context) error {

	articleId, err := strconv.Atoi(e.Param("articleId"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"err": err.Error(),
		})
	}

	userId, err := middlewares.VerifyAccessToken(e)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err": err.Error(),
		})
	}

	comment := request.CommentRequest{}
	e.Bind(&comment)

	err = uh.CommentBusiness.AddComment(comment.Content, articleId, userId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  "comment posted!",
	})
}

func (uh *CommentHandler) UpdateComment(e echo.Context) error {
	_, err := strconv.Atoi(e.Param("articleId"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err": err.Error(),
		})
	}

	commentId, err := strconv.Atoi(e.Param("commentId"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err": err.Error(),
		})
	}

	userId, err := middlewares.VerifyAccessToken(e)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err": err.Error(),
		})
	}

	comment := request.CommentRequest{}
	e.Bind(&comment)
	err = uh.CommentBusiness.UpdateComment(comment.Content, commentId, userId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  "comment updated!",
	})

}

func (uh *CommentHandler) DeleteComment(e echo.Context) error {
	_, err := strconv.Atoi(e.Param("articleId"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err": err.Error(),
		})
	}

	commentId, err := strconv.Atoi(e.Param("commentId"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err": err.Error(),
		})
	}

	userId, err := middlewares.VerifyAccessToken(e)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err": err.Error(),
		})
	}

	err = uh.CommentBusiness.DeleteComment(commentId, userId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "comment deleted",
	})

}

func (alh *CommentHandler) GetArticleComments(e echo.Context) error {
	articleId, err := strconv.Atoi(e.Param("articleId"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err":     err.Error(),
		})
	}
	data, err := alh.CommentBusiness.GetArticleComments(articleId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"data":   response.ToCommentResponseList(data),
	})

}
