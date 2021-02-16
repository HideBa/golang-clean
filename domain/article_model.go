package domain

type ArticleModel struct {
	ID      uint64
	Content string
	UserID  uint64
}

func NewArticleModel(content string, userID uint64) *ArticleModel {
	return &ArticleModel{Content: content, UserID: userID}
}
