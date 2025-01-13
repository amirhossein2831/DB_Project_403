package make

import (
	"DB_Project/cmd/make/controller"
	"DB_Project/cmd/make/exception"
	"github.com/spf13/cobra"
)

// Make Commands for making file
var Make = &cobra.Command{
	Use:   "make",
	Short: "Commands for making file",
}

func init() {
	Make.AddCommand(
		exception.ExceptionCmd,
		controller.ControllerCmd,
	)
}
