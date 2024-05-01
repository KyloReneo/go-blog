package services

//Business Logic layer on repository layer
//This is an intermediate between the Presentation Layer and the Data Access Layer
import (
	"github.com/kyloReneo/go-blog/internal/modules/article/requests/articles"
	ArticleResponse "github.com/kyloReneo/go-blog/internal/modules/article/responses"
	UserResponse "github.com/kyloReneo/go-blog/internal/modules/user/responses"
)

type ArticleServiceInterface interface {
	GetFeaturedArticles() ArticleResponse.Articles
	GetStoriesArticles() ArticleResponse.Articles
	Find(id int) (ArticleResponse.Article, error)
	StoreAsUser(request articles.StoreRequest, user UserResponse.User) (ArticleResponse.Article, error)
}
