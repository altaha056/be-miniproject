package response

import (
	"antonio/features/contents"
	"time"
)

type ArticleResponse struct {
	ID        int          `json:"id"`
	Title     string       `json:"title"`
	Image     string       `json:"image"`
	Content   string       `json:"content"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	Tools     []string     `json:"tools"`
	User      UserResponse `json:"user"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func ToArticleResponse(article contents.Core) ArticleResponse {
	return ArticleResponse{
		ID:        article.ID,
		Title:     article.Title,
		Image:     article.Image,
		Content:   article.Content,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
		User:      toUserResponse(article.User),
		Tools:     toTagsResponse(article.Tools),
	}
}

func toUserResponse(article contents.UserCore) UserResponse {
	return UserResponse{
		ID:       article.ID,
		Email:    article.Email,
		Username: article.Username,
	}
}

func toTagsResponse(tags []contents.ToolCore) []string {
	convertedTags := make([]string, 0, len(tags))
	for _, v := range tags {
		convertedTags = append(convertedTags, v.Title)
	}
	return convertedTags
}

func ToArticleResponseList(articleList []contents.Core) []ArticleResponse {
	convertedArticle := []ArticleResponse{}
	for _, article := range articleList {
		convertedArticle = append(convertedArticle, ToArticleResponse(article))
	}

	return convertedArticle
}
