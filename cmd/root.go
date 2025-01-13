package cmd

import (
	"DB_Project/cmd/app"
	"DB_Project/cmd/database"
	make2 "DB_Project/cmd/make"
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
		make2.Make,
	)
}

func Execute() error {
	return rootCmd.Execute()
}
