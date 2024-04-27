package controllers

import (
	"github.com/gin-gonic/gin"
)

// Define a controller type struct and a function that returns a Controller instance
type Controller struct {
}

func New() *Controller {
	return &Controller{}
}

// A handler function for GET the "/register" requestes
func (controller *Controller) Register(ctx *gin.Context) {

}

// A handler function for POST the "/register" requestes
func (controller *Controller) HandleRegister(ctx *gin.Context) {

}
