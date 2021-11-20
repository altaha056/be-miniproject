package response

import (
	"mpbe/features/posts"
	"time"
)

type Post struct {
	ID        int `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"updated_at"`
	UrlImg	string `json:"url_img"`
	Caption string `json:"caption"`
}

func FromCore(core posts.Core)Post{
	return Post{
		ID: core.ID,
		CreatedAt: core.CreatedAt,
		UpdateAt: core.UpdatedAt,
		UrlImg: core.UrlImg,
		Caption: core.Caption,
	}
}

func FromCoreSlice(core []posts.Core) []Post {
	var pstArray []Post
	for key := range core {
		pstArray = append(pstArray, FromCore(core[key]))
	}
	return pstArray
}