package repositories

//Abstraction layer between data access layer and Business Logic layer
import (
	"gorm.io/gorm"

	userModel "github.com/kyloReneo/go-blog/internal/modules/user/models"
	"github.com/kyloReneo/go-blog/pkg/database"

)

// Article Repository struct that is made of gorm.DB
type userRepository struct {
	DB *gorm.DB
}

// New() Method Returns a new database connection when is called
func New() *userRepository {
	return &userRepository{
		DB: database.Connection(),
	}
}

// Finds the requested id's article
func (userRepository *userRepository) Create(user userModel.User) userModel.User {
	var newUser userModel.User

	userRepository.DB.Create(&user).Scan(&newUser)

	return newUser
}
