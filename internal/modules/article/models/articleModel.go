package models

import (
	"gorm.io/gorm"

	"github.com/kyloReneo/go-blog/internal/modules/user/models"
)


//Article Data Access Layer
type Article struct {
	gorm.Model
	Title   string `gorm:"varchar:191"`
	Content string `gorm:"text"`
	UserID  uint
	User    models.User
}
