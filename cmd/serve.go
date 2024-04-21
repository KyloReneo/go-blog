package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/kyloReneo/go-blog/pkg/config"
	"github.com/kyloReneo/go-blog/pkg/routing"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serves app on dev server",
	Long:  `Application will be served on host and port defined in config.yml file.`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {

	//Set configs
	config.Set()

	//Initial gin router and return it
	routing.Init()
	router := routing.GetRouter()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"app name": viper.Get("App.Name"),
		})
	})
	routing.Serve()
}
