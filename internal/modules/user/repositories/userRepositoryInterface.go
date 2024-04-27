package repositories

//Abstraction layer between data access layer and service layer
import userModel "github.com/kyloReneo/go-blog/internal/modules/user/models"

type userRepositoryInterface interface {
	// Creates a user in database
	Create(user userModel.User) userModel.User
}
