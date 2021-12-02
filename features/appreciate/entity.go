package appreciate

import "time"

type Core struct {
	UserId    int
	User      UserCore
	ArticleId int
	Article   ArticleCore
}

type UserCore struct {
	ID       int
	Username string
	Email    string
}

type ArticleCore struct {
	ID        int
	Title     string
	Image     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserId    int
	User      UserCore
	Tools     []ToolCore
}

type ToolCore struct {
	ID    int
	Title string
}

type Business interface {
	Upvote(articleId, userId int) error
	Rating(userId int) ([]ArticleCore, error)
	WhoVote(articleId int) ([]UserCore, error)
	Downvote(articleId, userId int) error
}

type Data interface {
	Upvote(articleId, userId int) error
	Rating(userId int) ([]ArticleCore, error)
	WhoVote(articleId int) ([]UserCore, error)
	Downvote(articleId, userId int) error
}
