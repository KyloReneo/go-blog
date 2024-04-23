package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/kyloReneo/go-blog/pkg/config"
)

func Connect() {
	cfg := config.Get().DB

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Cannot connect to the database due to the original error:\n %v", err)
		return
	}

	DB = db
}
