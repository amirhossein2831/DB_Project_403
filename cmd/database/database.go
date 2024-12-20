package database

import (
	"DB_Project/cmd/database/migrate"
	"DB_Project/cmd/database/seed"
	"github.com/spf13/cobra"
)

// Database Commands for interacting with apps
var Database = &cobra.Command{
	Use:   "database",
	Short: "Commands for interacting with database.",
}

func init() {
	Database.AddCommand(
		migrate.Migrate,
		seed.Seed,
	)
}
