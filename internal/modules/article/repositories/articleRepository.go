package repositories

//Abstraction layer between data access layer and Business Logic layer
import (
	"gorm.io/gorm"

	articleModel "github.com/kyloReneo/go-blog/internal/modules/article/models"
	"github.com/kyloReneo/go-blog/pkg/database"
)

// Article Repository struct that is made of gorm.DB
type ArticleRepository struct {
	DB *gorm.DB
}

// New() Method Returns a new database connection when is called
func New() *ArticleRepository {
	return &ArticleRepository{
		DB: database.Connection(),
	}
}

// Gets an articleRepository struct as reciver and returns limit sized of articles from the article table
func (articleRepository *ArticleRepository) List(limit int) []articleModel.Article {
	var articles []articleModel.Article

	articleRepository.DB.Limit(limit).Joins("User").Order("rand()").Find(&articles)

	return articles
}

// Finds the requested id's article
func (articleRepository *ArticleRepository) Find(id int) articleModel.Article {
	var article articleModel.Article

	articleRepository.DB.Joins("User").Find(&article, id)

	return article
}

func (articleRepository *ArticleRepository) Create(article articleModel.Article) articleModel.Article {

	var newArticle articleModel.Article
	articleRepository.DB.Create(&article).Scan(&newArticle)

	return newArticle
}
