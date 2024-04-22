package routing

import (
	"fmt"
	"log"

	"github.com/kyloReneo/go-blog/pkg/config"
)

func Serve() {
	configs := config.Get()

	err := router.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port)) // listen and serve on 0.0.0.0:8080
	if err != nil {
		log.Fatalf("Something went wrong in routing. Error:\n %v", err)
	}

}
