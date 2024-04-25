package responses

import (
	"fmt"

	userModel "github.com/kyloReneo/go-blog/internal/modules/user/models"

)

// Response structure for the user responses
type User struct {
	ID    uint
	Image string
	Name  string
	Email string
}

// Structure for return bunch of users
type Users struct {
	Data []User
}

// Recives the model object and transforms it to response object
func ToUser(user userModel.User) User {
	return User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Image: fmt.Sprintf("https://ui-avatars.com/api/?name=%s", user.Name),
	}
}
