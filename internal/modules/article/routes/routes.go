package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/kyloReneo/go-blog/internal/middlewares"
	articleCtrl "github.com/kyloReneo/go-blog/internal/modules/article/controllers"
)

func Routes(router *gin.Engine) {

	articleController := articleCtrl.New()
	router.GET("/articles/:id", articleController.Show)

	authGroup := router.Group("/articles")
	authGroup.Use(middlewares.IsAuthenticated())
	{
		authGroup.GET("/create", articleController.Create)
		authGroup.POST("/store", articleController.Store)
	}

}
