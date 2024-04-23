package bootstrap

import (
	"github.com/kyloReneo/go-blog/pkg/config"
	"github.com/kyloReneo/go-blog/pkg/html"
	"github.com/kyloReneo/go-blog/pkg/routing"

)

func Serve() {
	config.Set()
	routing.Init()
	html.LoadHTML(routing.GetRouter())
	routing.RegisterRoutes()
	routing.Serve()
}
