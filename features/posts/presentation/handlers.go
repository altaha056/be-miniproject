package c_posts

import (
	"mpbe/features/posts"
	presention_response "mpbe/features/posts/presentation/response"
	"net/http"

	"github.com/labstack/echo"
)

type PostsHandler struct {
	postsBusiness posts.Bussiness
}

func NewPostHandler(pb posts.Bussiness)*PostsHandler{
	return &PostsHandler{
		postsBusiness: pb,
	}
}

func (ph*PostsHandler)GetAllPost(c echo.Context)error{
	result:=ph.postsBusiness.GetAllData("")

	return c.JSON(http.StatusOK,map[string]interface{}{
		"message": "everything work",
		"data": presention_response.FromCoreSlice(result),
	})
}