package services

//Business Logic layer on repository layer
//This is an intermediate between the Presentation Layer and the Data Access Layer
import ArticleModel "github.com/kyloReneo/go-blog/internal/modules/article/models"

type ArticleServiceInterface interface {
	GetFeaturedArticles() []ArticleModel.Article
	GetStoriesArticles() []ArticleModel.Article
}
