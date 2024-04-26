package repositories

//Abstraction layer between data access layer and service layer
import articleModel "github.com/kyloReneo/go-blog/internal/modules/article/models"

type ArticleRepositoryInterface interface {
	List(limit int) []articleModel.Article
	Find(id int) articleModel.Article
}
