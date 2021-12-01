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
		SigningKey: []byte(constants.ACCESS_TOKEN_KEY),
		Claims:     &middlewares.JwtCustomClaims{},
	}

	presenter := factory.Init()

	// eAuth := e.Group("/auth")
	// eAuth.PUT("", presenter.AuthHandler.ReLoginHandler)
	// eAuth.DELETE("", presenter.AuthHandler.LogoutHandler)


	e.POST("/user/login", presenter.AuthHandler.LoginHandler)
	e.POST("/user/register", presenter.UserHandler.RegisterUserHandler)


	e.GET("/user", presenter.UserHandler.GetAllUsersHandler)
	e.GET("/user/:userId", presenter.UserHandler.GetUserByIdHandler)
	e.GET("/user/:userId/likes", presenter.ArticleLikesHandler.GetLikedArticles)
	e.GET("/user/:userId/articles", presenter.ArticleHandler.GetAllUserArticlesHandler)


	e.POST("/content", presenter.ArticleHandler.CreateArticleHandler, middleware.JWTWithConfig(configJWT))
	e.GET("/content", presenter.ArticleHandler.GetAllArticleHandler)
	e.GET("/content/:articleId", presenter.ArticleHandler.GetArticleByIdHandler)
	e.PUT("/content/:articleId", presenter.ArticleHandler.UpdateArticleByIdHandler, middleware.JWTWithConfig(configJWT))
	e.DELETE("/content/:articleId", presenter.ArticleHandler.DeleteArticleByIdHandler, middleware.JWTWithConfig(configJWT))
	
	
	e.GET("/content/:articleId/likes", presenter.ArticleLikesHandler.GetLikingUsers)
	e.PUT("/content/:articleId/likes", presenter.ArticleLikesHandler.LikeArticle, middleware.JWTWithConfig(configJWT))
	e.DELETE("/content/:articleId/likes", presenter.ArticleLikesHandler.UnlikeArticle, middleware.JWTWithConfig(configJWT))
	

	e.GET("/content/:articleId/comments", presenter.CommentHandler.GetArticleComments)
	e.POST("/content/:articleId/comments", presenter.CommentHandler.AddComment)
	e.PATCH("/content/:articleId/comments/:commentId", presenter.CommentHandler.UpdateComment)
	e.DELETE("/content/:articleId/comments/:commentId", presenter.CommentHandler.DeleteComment)

	return e

}
