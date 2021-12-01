package factory

import (
	"antonio/db"

	// auth domain
	authBusiness "antonio/features/authentication/business"
	authData "antonio/features/authentication/data"
	authPresentation "antonio/features/authentication/presentation"

	// user domain
	userBusiness "antonio/features/users/business"
	userData "antonio/features/users/data"
	userPresentation "antonio/features/users/presentation"

	// article domain
	articleBusiness "antonio/features/articles/business"
	articleData "antonio/features/articles/data"
	articlePresentation "antonio/features/articles/presentation"

	// appreciate domain
	appreciateBusiness "antonio/features/appreciate/business"
	appreciateData "antonio/features/appreciate/data"
	appreciatePresentation "antonio/features/appreciate/presentation"

	// comments domain
	commentsBusiness "antonio/features/comments/business"
	commentsData "antonio/features/comments/data"
	commentsPresentation "antonio/features/comments/presentation"
)

type Presenter struct {
	AuthHandler         authPresentation.AuthHandler
	UserHandler         userPresentation.UserHandler
	ArticleHandler      articlePresentation.ArticleHandler
	RatingHandler    	appreciatePresentation.RatingHandler
	CommentHandler      commentsPresentation.CommentHandler
}

func Init() Presenter {
	// auth layer
	authData := authData.NewMysqlAuthRepository(db.DB)
	authBusiness := authBusiness.NewAuthBusiness(authData)
	authPresentation := authPresentation.NewAuthHandler(authBusiness)

	// users layer
	userData := userData.NewMysqlUserRepository(db.DB)
	userBusiness := userBusiness.NewUserBusiness(userData)
	userPresentation := userPresentation.NewUserHandler(userBusiness)

	// articles layer
	articleData := articleData.NewMysqlArticleRepository(db.DB)
	articleBusiness := articleBusiness.NewArticleBusiness(articleData)
	articlePresentation := articlePresentation.NewArticleHandler(articleBusiness)

	// article appreciate layer
	appreciateData := appreciateData.NewMysqlRatingRepository(db.DB)
	appreciateBusiness := appreciateBusiness.NewRatingBusiness(appreciateData)
	appreciatePresentation := appreciatePresentation.NewRatingHandler(appreciateBusiness)

	// comments layer
	commentsData := commentsData.NewMysqlCommentsRepository(db.DB)
	commentsBusiness := commentsBusiness.NewCommentsBusiness(commentsData)
	commentsPresentation := commentsPresentation.NewCommentHandler(commentsBusiness)

	return Presenter{
		AuthHandler:         *authPresentation,
		UserHandler:         *userPresentation,
		ArticleHandler:      *articlePresentation,
		RatingHandler: 	 	 *appreciatePresentation,
		CommentHandler:      *commentsPresentation,
	}
}
