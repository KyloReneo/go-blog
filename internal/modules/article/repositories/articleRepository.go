package repositories

import (
	"gorm.io/gorm"

	articleModel "github.com/kyloReneo/go-blog/internal/modules/article/models"
	"github.com/kyloReneo/go-blog/pkg/database"

)

type ArticleRepository struct {
	DB *gorm.DB
}

func New() *ArticleRepository {
	return &ArticleRepository{
		DB: database.Connection(),
	}
}

func (articleRepository *ArticleRepository) List(limit int) []articleModel.Article {
	var articles []articleModel.Article

	articleRepository.DB.Limit(limit).Joins("User").Order("rand()").Find(&articles)

	return articles
}
