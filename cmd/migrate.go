package cmd

import (
	"github.com/spf13/cobra"

	"github.com/kyloReneo/go-blog/pkg/bootstrap"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Table migration",
	Long:  `Migrates tables.`,
	Run: func(cmd *cobra.Command, args []string) {
		Migrate()
	},
}

func Migrate() {
	bootstrap.Migrate()
}
