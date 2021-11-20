package factory

import (
	_post_business "mbpe/features/posts/bussiness"
	_post_data "mpbe/features/posts/data"

	_post_presentation "mbpe/features/posts/presentation"
	"mpbe/config"
)

type Presenter struct {
	PostPresentation *_post_presentation.PostsHandler
}

func Init() Presenter {

	postData := _post_data.NewPostRepository(config.DB)
	postBussiness := _post_business.NewAricleBussiness(postData)
	postPresentation := _post_presentation.NewPostHandler(postBussiness)

	return Presenter{
		PostPresentation: postPresentation,
	}
}
