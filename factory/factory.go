package factory

import (
	"antonio/db"

	// auth entity
	authBusiness "antonio/features/authentication/business"
	authData "antonio/features/authentication/data"
	authPresentation "antonio/features/authentication/presentation"

	// user entity
	userBusiness "antonio/features/users/business"
	userData "antonio/features/users/data"
	userPresentation "antonio/features/users/presentation"

	// content entity
	articleBusiness "antonio/features/contents/business"
	articleData "antonio/features/contents/data"
	articlePresentation "antonio/features/contents/presentation"

	// rating entity
	appreciateBusiness "antonio/features/appreciate/business"
	appreciateData "antonio/features/appreciate/data"
	appreciatePresentation "antonio/features/appreciate/presentation"

	// comments entity
	commentsBusiness "antonio/features/comments/business"
	commentsData "antonio/features/comments/data"
	commentsPresentation "antonio/features/comments/presentation"

	//news entitiy
	newsData "antonio/features/third-party-news/data"
	newsPresent "antonio/features/third-party-news/presentation"
	newsService "antonio/features/third-party-news/business"
)

type Presenter struct {
	AuthHandler         authPresentation.AuthHandler
	UserHandler         userPresentation.UserHandler
	ArticleHandler      articlePresentation.ArticleHandler
	RatingHandler    	appreciatePresentation.RatingHandler
	CommentHandler      commentsPresentation.CommentHandler
	NewsPresentation    newsPresent.NewsHandler
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

	// content layer
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

	// news 3rd party api layer
	newsData := newsData.NewNewsApiRepository("http://api.mediastack.com/v1/news", "dd327eaab090f4964b6370f78d3591f3")
	newsService := newsService.NewApiService(newsData)

	
	return Presenter{
		AuthHandler:         *authPresentation,
		UserHandler:         *userPresentation,
		ArticleHandler:      *articlePresentation,
		RatingHandler: 	 	 *appreciatePresentation,
		CommentHandler:      *commentsPresentation,
		NewsPresentation:	 *newsPresent.NewNewsHandler(newsService),
	}
}
