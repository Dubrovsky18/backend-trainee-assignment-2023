package cli

import (
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/pkg/logger"
	"github.com/spf13/cobra"
)

// ExecuteRootCmd prepares all CLI commands
func ExecuteRootCmd() {
	c := cobra.Command{}

	c.AddCommand(NewServeCmd())

	if err := c.Execute(); err != nil {
		logger.Fatal(err.Error())
	}
}
