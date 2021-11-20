package factory

import (
	_post_business "mpbe/features/posts/business"
	_post_data "mpbe/features/posts/data"

	_post_presentation "mpbe/features/posts/presentation"
	"mpbe/config"
)

type Presenter struct {
	PostPresentation *_post_presentation.PostsHandler
}

func Init() Presenter {

	postData:=_post_data.NewPostRepository(config.DB)
	postBusiness:=_post_business.NewPostBusiness(postData)
	postPresentation:=_post_presentation.NewPostHandler(postBusiness)

	return Presenter{
		PostPresentation: postPresentation,
	}
}
