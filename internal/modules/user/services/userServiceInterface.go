package services

//Business Logic layer on repository layer
//This is an intermediate between the Presentation Layer and the Data Access Layer
import (
	"github.com/kyloReneo/go-blog/internal/modules/user/requests/auth"
	UserResponse "github.com/kyloReneo/go-blog/internal/modules/user/responses"
)

type UserServiceInterface interface {
	Create(request auth.RegisterRequest) (UserResponse.User, error)
	CheckUserExists(email string) bool
	HandleUsersLogin(request auth.LoginRequest) (UserResponse.User, error)
}
