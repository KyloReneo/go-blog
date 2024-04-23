package cmd

import (
	"github.com/spf13/cobra"

	"github.com/kyloReneo/go-blog/pkg/bootstrap"

)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serves app on dev server",
	Long:  `Application will be served on host and port defined in config.yml file.`,
	Run: func(cmd *cobra.Command, args []string) {
		Serve()
	},
}

func Serve() {
	bootstrap.Serve()
}
