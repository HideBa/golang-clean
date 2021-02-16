package domain

type ArticleRepository interface {
	CreateArticle(newArticle *ArticleModel) (*ArticleModel, error)
	GetArticleById(id uint64) (*ArticleModel, error)
	GetArticleByUserID(userID uint64) ([]*ArticleModel, error)
	UpdateArticle(newArticle *ArticleModel) error
	DeleteArticle(id uint64) error
}
