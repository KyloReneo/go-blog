package bootstrap

import (
	"github.com/kyloReneo/go-blog/pkg/config"
	"github.com/kyloReneo/go-blog/pkg/routing"
)

func Serve() {
	config.Set()
	routing.Init()
	routing.RegisterRoutes()
	routing.Serve()
}
