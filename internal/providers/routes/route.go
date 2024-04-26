package routes

import (
	"github.com/gin-gonic/gin"

	articleRoutes "github.com/kyloReneo/go-blog/internal/modules/article/routes"
	homeRoutes "github.com/kyloReneo/go-blog/internal/modules/home/routes"
)

func RegisterRoutes(router *gin.Engine) {
	homeRoutes.Routes(router)
	articleRoutes.Routes(router)
}
