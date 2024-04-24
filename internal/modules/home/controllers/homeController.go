package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	ArticleRepository "github.com/kyloReneo/go-blog/internal/modules/article/repositories"
	"github.com/kyloReneo/go-blog/pkg/html"

)

// Define a controller type struct and a function that returns a Controller instance
type Controller struct {
	articleRepository ArticleRepository.ArticleRepositoryInterface
}

func New() *Controller {
	return &Controller{
		articleRepository: ArticleRepository.New(),
	}
}

// Create a handler function for returning the gin context
func (controller *Controller) Index(ctx *gin.Context) {

	html.Render(ctx, http.StatusOK, "../html/home", gin.H{
		"title": "Home Page",
	})

}
