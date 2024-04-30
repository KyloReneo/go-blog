package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/kyloReneo/go-blog/internal/modules/user/requests/auth"
	userService "github.com/kyloReneo/go-blog/internal/modules/user/services"
	"github.com/kyloReneo/go-blog/pkg/converters"
	"github.com/kyloReneo/go-blog/pkg/errors"
	"github.com/kyloReneo/go-blog/pkg/html"
	"github.com/kyloReneo/go-blog/pkg/old"
	"github.com/kyloReneo/go-blog/pkg/sessions"

)

// Define a controller type struct and a function that returns a Controller instance
type Controller struct {
	userSercive userService.UserServiceInterface
}

func New() *Controller {
	return &Controller{
		userSercive: userService.New(),
	}
}

// A handler function for GET the "/register" requestes
func (controller *Controller) Register(ctx *gin.Context) {

	html.Render(ctx, http.StatusOK, "../../user/html/register", gin.H{
		"title": "Register",
	})
}

// A handler function for POST the "/register" requestes
func (controller *Controller) HandleRegister(ctx *gin.Context) {

	// Validate the request
	var request auth.RegisterRequest

	// This will infer what binder to use depending on the content-type header.
	if err := ctx.ShouldBind(&request); err != nil {

		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(ctx, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(ctx)
		sessions.Set(ctx, "old", converters.UrlValuesToString(old.Get()))

		ctx.Redirect(http.StatusFound, "/register")
		return
	}

	if controller.userSercive.CheckUserExists(request.Email) {
		errors.Init()
		errors.Add("Email", "Email address already exists")
		sessions.Set(ctx, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(ctx)
		sessions.Set(ctx, "old", converters.UrlValuesToString(old.Get()))

		ctx.Redirect(http.StatusFound, "/register")
		return
	}

	// Create the user
	user, err := controller.userSercive.Create(request)

	// Check if there is any error on the user creation
	if err != nil {
		ctx.Redirect(http.StatusFound, "/register")
		return
	}

	sessions.Set(ctx, "auth", strconv.Itoa(int(user.ID)))

	// After creating the user redirects to the home page
	log.Printf("The user with %s Email created successfully.\n", user.Email)
	ctx.Redirect(http.StatusFound, "/")
}

// A handler function for GET the "/login" requestes
func (controller *Controller) Login(ctx *gin.Context) {

	html.Render(ctx, http.StatusOK, "../../user/html/login", gin.H{
		"title": "Login",
	})
}

// A handler function for POST the "/login" requestes
func (controller *Controller) HandleLogin(ctx *gin.Context) {

	// Validate the request
	// Create an instance of binded login request
	var request auth.LoginRequest

	// This will infer what binder to use depending on the content-type header.
	if err := ctx.ShouldBind(&request); err != nil {

		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(ctx, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(ctx)
		sessions.Set(ctx, "old", converters.UrlValuesToString(old.Get()))

		ctx.Redirect(http.StatusFound, "/login")
		return
	}

	user, err := controller.userSercive.HandleUsersLogin(request)
	if err != nil {
		errors.Init()
		errors.Add("email", err.Error())
		sessions.Set(ctx, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(ctx)
		sessions.Set(ctx, "old", converters.UrlValuesToString(old.Get()))

		ctx.Redirect(http.StatusFound, "/login")
		return
	}
	sessions.Set(ctx, "auth", strconv.Itoa(int(user.ID)))

	// After login the user redirects to the home page
	log.Printf("The user with %s Email logded in successfully.\n", user.Email)
	ctx.Redirect(http.StatusFound, "/")
}
