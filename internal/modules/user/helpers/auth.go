package helpers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	userRepository "github.com/kyloReneo/go-blog/internal/modules/user/repositories"
	userResponse "github.com/kyloReneo/go-blog/internal/modules/user/responses"
	"github.com/kyloReneo/go-blog/pkg/sessions"

)

func Auth(ctx *gin.Context) userResponse.User {

	var response userResponse.User

	authID := sessions.Get(ctx, "auth")
	userID, _ := strconv.Atoi(authID)
	userRepo := userRepository.New()
	user := userRepo.FindByID(uint(userID))

	if user.ID == 0 {
		return response
	}

	return userResponse.ToUser(user)
}
