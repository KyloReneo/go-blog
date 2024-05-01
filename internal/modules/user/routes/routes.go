package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/kyloReneo/go-blog/internal/middlewares"
	usrCtrl "github.com/kyloReneo/go-blog/internal/modules/user/controllers"
)

func Routes(router *gin.Engine) {

	userController := usrCtrl.New()

	// Define guest group valid routes
	guestGroup := router.Group("/")
	guestGroup.Use(middlewares.IsGuest())
	{
		guestGroup.GET("/register", userController.Register)
		guestGroup.POST("/register", userController.HandleRegister)

		guestGroup.GET("/login", userController.Login)
		guestGroup.POST("/login", userController.HandleLogin)

	}

	// Define authenticated group accessable routes
	authGroup := router.Group("/")
	authGroup.Use(middlewares.IsAuthenticated())
	{
		router.POST("/logout", userController.HandleLogout)
	}

}
