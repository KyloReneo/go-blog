package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	ArticleService "github.com/kyloReneo/go-blog/internal/modules/article/services"
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
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"Error": "Somthing went wrong with converting the id.",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Message": id,
		})
	}
	//Find the article in the database

	//Show error page if the article is not found

	//render the article template

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "controller is Ok.",
	})

}
