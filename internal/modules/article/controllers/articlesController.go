package controllers

import (
	"net/http"
	"strconv"

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

// Creates a handler function for returning the gin context
func (controller *Controller) Show(ctx *gin.Context) {

	//Get the article id
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		html.Render(ctx,
			http.StatusInternalServerError,
			"/templates/errors/html/500",
			gin.H{
				"title":   "Internal error",
				"Message": "Converting id error",
			})
		return
	}

	//Find the article in the database
	article, err := controller.articleService.Find(id)
	//Show error page if the article is not found
	if err != nil {
		html.Render(ctx,
			http.StatusNotFound,
			"/templates/errors/html/404",
			gin.H{
				"title":   "Page not found",
				"Message": err.Error(),
			})
		return

	}

	//render the article template
	ctx.JSON(http.StatusOK, gin.H{
		"article": article,
	})

}
