package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kyloReneo/go-blog/pkg/html"

)

// Define a controller type struct and a function that returns a Controller instance
type Controller struct {
}

func New() *Controller {
	return &Controller{}
}

// A handler function for GET the "/register" requestes
func (controller *Controller) Register(ctx *gin.Context) {

	html.Render(ctx, http.StatusOK,"../../user/html/register", gin.H{
		
	})
}

// A handler function for POST the "/register" requestes
func (controller *Controller) HandleRegister(ctx *gin.Context) {

}
