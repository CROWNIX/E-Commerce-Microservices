package workers

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/CROWNIX/E-Commerce-Microservices/api-gateway/internal/config"
	"github.com/CROWNIX/E-Commerce-Microservices/api-gateway/internal/log"
	"github.com/CROWNIX/E-Commerce-Microservices/api-gateway/internal/provider"
	"github.com/spf13/cobra"
)

func Start() {
	err := config.LoadConfig()
	if err != nil{
		log.Logger.Fatal(err)
	}

	provider.BootstrapGlobal()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	rootCmd := &cobra.Command{}
	cmd := []*cobra.Command{
		{
			Use:   "serve-all",
			Short: "Run all",
			Run: func(cmd *cobra.Command, _ []string) {
				runHttpWorker(config.GetConfig(), ctx)
			},
		},
	}

	rootCmd.AddCommand(cmd...)
	if err := rootCmd.Execute(); err != nil {
		log.Logger.Fatal(err)
	}
}
