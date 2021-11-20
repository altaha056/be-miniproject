package factory

import (
	_post_business "mbpe/features/articles/bussiness"
	_post_data "mpbe/features/articles/data"

	_post_presentation "mbpe/features/articles/presentation"
	"mpbe/config"
)

type Presenter struct {
	ArticlePresentation *_article_presentation.ArticlesHandler
}

func Init() Presenter {

	postData := _post_data.NewArticleRepository(config.DB)
	postBussiness := _post_business.NewAricleBussiness(postData)
	postPresentation := _post_presentation.NewArticleHandler(postBussiness)

	return Presenter{
		ArticlePresentation: postPresentation,
	}
}
