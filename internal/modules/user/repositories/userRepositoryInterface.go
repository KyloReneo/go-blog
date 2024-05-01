package repositories

//Abstraction layer between data access layer and service layer
import userModel "github.com/kyloReneo/go-blog/internal/modules/user/models"

type UserRepositoryInterface interface {
	// Creates a user in database
	Create(user userModel.User) userModel.User
	FindByEmail(email string) userModel.User
	FindByID(id int) userModel.User
}
