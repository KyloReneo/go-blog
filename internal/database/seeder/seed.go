package seeder

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"

	articleModel "github.com/kyloReneo/go-blog/internal/modules/article/models"
	userModel "github.com/kyloReneo/go-blog/internal/modules/user/models"
	"github.com/kyloReneo/go-blog/pkg/database"

)

// Seeder logic
func Seed() {
	db := database.Connection()

	//Generate some hashed password
	password := []byte("Secert")
	hasedPassword, err := bcrypt.GenerateFromPassword(password, 12)
	if err != nil {
		log.Fatalf("Cannot hash the password due to the original Error:\n %v", err)
		return
	}

	//Create record based on its model
	user := userModel.User{Name: "Random name", Email: "Random email", Password: string(hasedPassword)}

	db.Create(&user) // pass pointer of data to Create
	id := user.ID    //returns created users userID

	fmt.Printf("User with %s email created successfully with %d UserID.", user.Email, id)

	//Create article seeds for created user
	for i := 1; i <= 10; i++ {
		article := articleModel.Article{Title: fmt.Sprintf("Random title %d ", i), Content: fmt.Sprintf("Random content %d ", i), UserID: id}
		db.Create(&article)

		fmt.Printf("Article %s was created by %d successfully.\n", article.Title, id)
	}

	log.Println("Seeder Done!")
}
