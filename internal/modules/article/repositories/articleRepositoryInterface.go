package repositories

import articleModel "github.com/kyloReneo/go-blog/internal/modules/article/models"

type ArticleRepositoryInterface interface {
	List(limit int) []articleModel.Article
}
