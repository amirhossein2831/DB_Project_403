package cmd

import (
	"DB_Project/cmd/app"
	"DB_Project/cmd/database"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "db-project",
	Short: "db-project CLI",
	Long:  "Golang db-project CLI",
}

func init() {
	rootCmd.AddCommand(
		app.App,
		database.Database,
	)
}

func Execute() error {
	return rootCmd.Execute()
}
