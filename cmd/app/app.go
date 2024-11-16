package app

import (
	"DB_Project/src/bootstrap"
	"github.com/spf13/cobra"
	"log"
	"time"
)

// App Commands for interacting with apps
var App = &cobra.Command{
	Use:   "app",
	Short: "Commands for interacting with apps.",
}

func init() {
	App.AddCommand(
		bootstrapCmd,
	)
}

// bootstrapCmd is sub command for App that bootstrap the application
// usage: go run main.go app bootstrap or ./binary app bootstrap
var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Bootstraps the application and it's related services.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Starting the Application...    " + "timestamp" + time.Now().String())
		bootstrap.Init()
		log.Println("Application exited successfully.    " + "timestamp" + time.Now().String())
	},
}
