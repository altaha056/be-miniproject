package response

import (
	"antonio/features/appreciate"
	"time"
)

type ArticleResponse struct {
	ID        int          `json:"id"`
	Title     string       `json:"title"`
	Image     string       `json:"image"`
	Content   string       `json:"content"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	Tools      []string    `json:"tools"`
	User      UserResponse `json:"user"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"Username"`
	Email    string `json:"email"`
}

func ToArticleResponse(article appreciate.ArticleCore) ArticleResponse {
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

func toUserResponse(user appreciate.UserCore) UserResponse {
	return UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
}

func toTagsResponse(tools []appreciate.ToolCore) []string {
	convertedTags := make([]string, 0, len(tools))
	for _, v := range tools {
		convertedTags = append(convertedTags, v.Title)
	}
	return convertedTags
}

func ToArticleResponseList(articleList []appreciate.ArticleCore) []ArticleResponse {
	convertedArticle := []ArticleResponse{}
	for _, article := range articleList {
		convertedArticle = append(convertedArticle, ToArticleResponse(article))
	}

	return convertedArticle
}

func ToUserResponseList(userList []appreciate.UserCore) []UserResponse {
	convertedArticle := []UserResponse{}
	for _, user := range userList {
		convertedArticle = append(convertedArticle, toUserResponse(user))
	}

	return convertedArticle
}
