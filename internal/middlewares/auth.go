package middlewares

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	userRepository "github.com/kyloReneo/go-blog/internal/modules/user/repositories"
	"github.com/kyloReneo/go-blog/pkg/sessions"

)

func IsAuthenticated() gin.HandlerFunc {

	var userRepo = userRepository.New()
	return func(ctx *gin.Context) {

		// If the user is not logged in redirect to the login page
		authID := sessions.Get(ctx, "auth")
		userID, _ := strconv.Atoi(authID)

		user := userRepo.FindByID(userID)
		if user.ID == 0 {
			ctx.Redirect(http.StatusFound, "/login")
			return
		}

		// If user is authenticated let user go to the next handler
		ctx.Next()

	}
}
