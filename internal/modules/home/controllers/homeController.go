package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	ArticleService "github.com/kyloReneo/go-blog/internal/modules/article/services"
	"github.com/kyloReneo/go-blog/pkg/html"

)

// Define a controller type struct and a function that returns a Controller instance
type Controller struct {
	articleService ArticleService.ArticleServiceInterface
}

func New() *Controller {
	return &Controller{
		articleService: ArticleService.New(),
	}
}

// Create a handler function for returning the gin context
func (controller *Controller) Index(ctx *gin.Context) {

	html.Render(ctx, http.StatusOK, "../html/home", gin.H{
		"title": "Home Page",
		"featured": controller.articleService.GetFeaturedArticles(),
		"stories":  controller.articleService.GetStoriesArticles(),
	})

	

}
