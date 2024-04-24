package services

//Service layer on repository layer
import ArticleModel "github.com/kyloReneo/go-blog/internal/modules/article/models"

type ArticleServiceInterface interface {
	GetFeaturedArticles() []ArticleModel.Article
	GetStoriesArticles() []ArticleModel.Article
}
