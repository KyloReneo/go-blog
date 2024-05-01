package repositories

//Abstraction layer between data access layer and Business Logic layer
import (
	"gorm.io/gorm"

	userModel "github.com/kyloReneo/go-blog/internal/modules/user/models"
	"github.com/kyloReneo/go-blog/pkg/database"

)

// Article Repository struct that is made of gorm.DB
type UserRepository struct {
	DB *gorm.DB
}

// New() Method Returns a new database connection when is called
func New() *UserRepository {
	return &UserRepository{
		DB: database.Connection(),
	}
}

// Creates and returns the new user
func (userRepository *UserRepository) Create(user userModel.User) userModel.User {
	var newUser userModel.User

	userRepository.DB.Create(&user).Scan(&newUser)

	return newUser
}

// Finds the user by email and returns the record as a user model response
func (userRepository *UserRepository) FindByEmail(email string) userModel.User {
	var user userModel.User

	userRepository.DB.First(&user, "email = ?", email)

	return user
}

// Finds the logged in user by id and returns the record as a user model response
func (userRepository *UserRepository) FindByID(id int) userModel.User {
	var user userModel.User

	userRepository.DB.First(&user, "id = ?", id)

	return user
}
