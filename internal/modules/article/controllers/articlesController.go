package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/kyloReneo/go-blog/internal/modules/article/requests/articles"
	ArticleService "github.com/kyloReneo/go-blog/internal/modules/article/services"
	"github.com/kyloReneo/go-blog/internal/modules/user/helpers"
	"github.com/kyloReneo/go-blog/pkg/converters"
	"github.com/kyloReneo/go-blog/pkg/errors"
	"github.com/kyloReneo/go-blog/pkg/html"
	"github.com/kyloReneo/go-blog/pkg/old"
	"github.com/kyloReneo/go-blog/pkg/sessions"
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
			"../../../templates/errors/html/500",
			gin.H{
				"title":   "Internal error",
				"message": "Converting id error",
			})
		return
	}

	//Find the article in the database
	article, err := controller.articleService.Find(id)
	//Show error page if the article is not found
	if err != nil {
		html.Render(ctx,
			http.StatusNotFound,
			"../../../templates/errors/html/404",
			gin.H{
				"title":   "Page not found",
				"message": err.Error(),
			})
		return

	}

	html.Render(ctx,
		http.StatusOK,
		"../../article/html/show",
		gin.H{
			"title":   "Show article",
			"article": article,
		})

}

// Handler function for creating articles
func (controller *Controller) Create(ctx *gin.Context) {
	html.Render(ctx,
		http.StatusOK,
		"../../article/html/create",
		gin.H{
			"title": "Create article",
		})
}

// Handler function for storing articles
func (controller *Controller) Store(ctx *gin.Context) {
	// Validate the request
	var request articles.StoreRequest

	// This will infer what binder to use depending on the content-type header.
	if err := ctx.ShouldBind(&request); err != nil {

		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(ctx, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(ctx)
		sessions.Set(ctx, "old", converters.UrlValuesToString(old.Get()))

		ctx.Redirect(http.StatusFound, "/articles/create")
		return
	}

	user := helpers.Auth(ctx)

	// Create the article
	article, err := controller.articleService.StoreAsUser(request, user)

	// Check if there is any error on the article creation
	if err != nil {
		ctx.Redirect(http.StatusFound, "/articles/create")
		return
	}

	ctx.Redirect(http.StatusFound, fmt.Sprintf("/articles/%d", article.ID))
}
