package routes

import (
	"github.com/gin-gonic/gin"

	articleCtrl "github.com/kyloReneo/go-blog/internal/modules/article/controllers"

)

func Routes(router *gin.Engine) {

	articleController := articleCtrl.New()
	router.GET("/articles/:id", articleController.Show)

}
