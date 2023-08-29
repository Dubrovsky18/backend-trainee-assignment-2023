package cli

import (
	"context"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/app"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/pkg/logger"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

// NewServeCmd starts new application instance
func NewServeCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "serve",
		Aliases: []string{"s"},
		Short:   "Start server",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info("starting application")

			sigchan := make(chan os.Signal, 1)
			signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			application, err := app.InitializeApplication()
			if err != nil {
				logger.Fatal("Failed initialize application", "error", err)
			}

			cliMode := false
			application.Start(ctx, cliMode)

			logger.Info("started")
			<-sigchan

			logger.Info("stop application", "error", application.Stop())

			time.Sleep(time.Second * cliCmdExecFinishDelaySeconds)
			logger.Info("finished")
		},
	}
}
