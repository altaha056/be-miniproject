package routes

import (
	"antonio/constants"
	"antonio/factory"
	"antonio/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	configJWT := middleware.JWTConfig{
		SigningKey: []byte(constants.ANOTHER_PASS),
		Claims:     &middlewares.JwtCustomClaims{},
	}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	
	presenter := factory.Init()

	e.POST("/user/login", presenter.AuthHandler.LoginHandler)
	e.POST("/user/register", presenter.UserHandler.RegisterUserHandler)


	e.GET("/user", presenter.UserHandler.GetAllUsersHandler)
	e.GET("/user/:userId", presenter.UserHandler.GetUserByIdHandler)
	e.GET("/user/:userId/rating", presenter.RatingHandler.Rating)
	e.GET("/user/:userId/content", presenter.ArticleHandler.GetAllUserArticlesHandler)


	e.POST("/content", presenter.ArticleHandler.CreateArticleHandler, middleware.JWTWithConfig(configJWT))
	e.GET("/content", presenter.ArticleHandler.GetAllArticleHandler)
	e.GET("/content/:contentId", presenter.ArticleHandler.GetArticleByIdHandler)
	e.PUT("/content/:contentId", presenter.ArticleHandler.UpdateArticleByIdHandler, middleware.JWTWithConfig(configJWT))
	e.DELETE("/content/:contentId", presenter.ArticleHandler.DeleteArticleByIdHandler, middleware.JWTWithConfig(configJWT))
	
	
	e.GET("/content/:contentId/rating", presenter.RatingHandler.WhoVote)
	e.PUT("/content/:contentId/upvote", presenter.RatingHandler.UpVote, middleware.JWTWithConfig(configJWT))
	e.DELETE("/content/:contentId/downvote", presenter.RatingHandler.DownVote, middleware.JWTWithConfig(configJWT))
	

	e.GET("/content/:contentId/comments", presenter.CommentHandler.GetArticleComments)
	e.POST("/content/:contentId/comments", presenter.CommentHandler.AddComment)
	e.PATCH("/content/:contentId/comments/:commentId", presenter.CommentHandler.UpdateComment)
	e.DELETE("/content/:contentId/comments/:commentId", presenter.CommentHandler.DeleteComment)

	e.GET("/news", presenter.NewsPresentation.GetNewsHandler)

	return e

}
