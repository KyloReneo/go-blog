package migration

import (
	"fmt"
	"log"

	articleModels "github.com/kyloReneo/go-blog/internal/modules/article/models"
	userModels "github.com/kyloReneo/go-blog/internal/modules/user/models"
	"github.com/kyloReneo/go-blog/pkg/database"
)

func Migrate() {

	db := database.Connection()
	err := db.AutoMigrate(&userModels.User{}, &articleModels.Article{})

	if err != nil {
		log.Fatalf("Migration failed due to the original Error:\n %v", err)
		return
	}
	fmt.Println("Migration Done!")
}
