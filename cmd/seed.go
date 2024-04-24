package cmd

import (
	"github.com/spf13/cobra"

	"github.com/kyloReneo/go-blog/pkg/bootstrap"

)

func init() {
	rootCmd.AddCommand(seedCmd)
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Database seeder",
	Long:  `Creates random seed for database table context.`,
	Run: func(cmd *cobra.Command, args []string) {
		Seed()
	},
}

func Seed() {
	bootstrap.Seed()
}
