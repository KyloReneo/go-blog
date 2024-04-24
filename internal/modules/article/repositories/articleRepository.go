package repositories

import (
	"gorm.io/gorm"

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
