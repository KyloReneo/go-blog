package services

//Business Logic layer on repository layer
//This is an intermediate between the Presentation Layer and the Data Access Layer
import (
	ArticleResponse "github.com/kyloReneo/go-blog/internal/modules/article/responses"

)

type ArticleServiceInterface interface {
	GetFeaturedArticles() ArticleResponse.Articles
	GetStoriesArticles() ArticleResponse.Articles
	//Find(id int) (ArticleResponse.Article, error)
}
