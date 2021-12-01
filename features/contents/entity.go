package contents

import "time"

type Core struct {
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

type UserCore struct {
	ID       int
	Email    string
	Username string
	Image    string
}

type ToolCore struct {
	ID    int
	Title string
}

type Business interface {
	CreateArticle(data Core, userId int) error
	GetAllArticles() ([]Core, error)
	GetArticleById(articleId int) (Core, error)
	UpdateArticleById(articleId int, data Core, userId int) error
	DeleteArticleById(articleId int, userId int) error
	GetAllUserArticles(userId int) ([]Core, error)
}

type Data interface {
	CreateTags(tags []ToolCore) ([]ToolCore, error)
	CreateArticle(data Core, userId int, tags []ToolCore) error
	GetAllArticles() ([]Core, error)
	GetArticleById(articleId int) (Core, error)
	UpdateArticleById(articleId int, data Core) error
	DeleteArticleById(articleId int) error
	VerifyArticleOwner(articleId int, userId int) error
	GetAllUserArticles(userId int) ([]Core, error)
}
