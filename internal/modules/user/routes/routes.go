package routes

import (
	"github.com/gin-gonic/gin"

	usrCtrl "github.com/kyloReneo/go-blog/internal/modules/user/controllers"
)

func Routes(router *gin.Engine) {

	userController := usrCtrl.New()
	router.GET("/register", userController.Register)
	router.POST("/register", userController.HandleRegister)

	router.GET("/login", userController.Login)
	router.POST("/login", userController.HandleLogin)
}
