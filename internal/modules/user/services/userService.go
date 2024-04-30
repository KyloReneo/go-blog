package services

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	userModel "github.com/kyloReneo/go-blog/internal/modules/user/models"
	userRepository "github.com/kyloReneo/go-blog/internal/modules/user/repositories"
	"github.com/kyloReneo/go-blog/internal/modules/user/requests/auth"
	userResponse "github.com/kyloReneo/go-blog/internal/modules/user/responses"
)

type UserSercive struct {
	userRepository userRepository.UserRepositoryInterface
}

func New() *UserSercive {
	return &UserSercive{
		userRepository: userRepository.New(),
	}
}

func (userService *UserSercive) Create(request auth.RegisterRequest) (userResponse.User, error) {
	var response userResponse.User
	var user userModel.User

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
	if err != nil {
		return response, errors.New("error hashing the password")
	}

	user.Name = request.Name
	user.Email = request.Email
	user.Password = string(hasedPassword)

	newUser := userService.userRepository.Create(user)
	if newUser.ID == 0 {
		return response, errors.New("error on creating new user")
	}
	return userResponse.ToUser(newUser), nil

}

func (userService *UserSercive) CheckUserExists(email string) bool {
	user := userService.userRepository.FindByEmail(email)

	return user.ID != 0
}

func (userService *UserSercive) HandleUsersLogin(request auth.LoginRequest) (userResponse.User, error) {

	var response userResponse.User
	existsUser := userService.userRepository.FindByEmail(request.Email)

	// Check if the inserted email is valid
	if existsUser.ID == 0 {
		return response, errors.New("invalid credentials")
	}

	// Check if the inserted password is vaid
	err := bcrypt.CompareHashAndPassword([]byte(existsUser.Password), []byte(request.Password))
	if err != nil {
		return response, errors.New("nvalid credentials")
	}

	return userResponse.ToUser(existsUser), nil
}
