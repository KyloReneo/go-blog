package bootstrap

import (
	"github.com/kyloReneo/go-blog/pkg/config"
	"github.com/kyloReneo/go-blog/pkg/database"
	"github.com/kyloReneo/go-blog/pkg/html"
	"github.com/kyloReneo/go-blog/pkg/routing"
	"github.com/kyloReneo/go-blog/pkg/static"

)

func Serve() {
	config.Set()
	database.Connect()
	routing.Init()
	static.LoadStatic(routing.GetRouter())
	html.LoadHTML(routing.GetRouter())
	routing.RegisterRoutes()
	routing.Serve()
}
