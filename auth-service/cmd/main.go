package main

import (
	"context"
	"log/slog"

	"auth-service/internal/config"
	"auth-service/internal/infra"
	"auth-service/internal/presentations"
	"pkg/telemetry"

	"github.com/CROWNIX/go-utils/validatorx"
	"github.com/spf13/cobra"
)

var restApiCmd = &cobra.Command{
	Use:  "rest-api",
	Long: "Rest API command",
	Run: func(cmd *cobra.Command, args []string) {
		if err := config.LoadConfig(); err != nil {
			slog.Error("failed to load config", "error", err)
			return
		}

		cfg := config.GetConfig()
		otelCleanup, err := telemetry.Init(context.Background(), telemetry.Config{
			Enabled:      cfg.Otel.Enabled,
			ServiceName:  cfg.AppName,
			Env:          cfg.AppEnv,
			OtlpEndpoint: cfg.Otel.Endpoint,
			OtlpUser:     cfg.Otel.Username,
			OtlpPassword: cfg.Otel.Password,
		})
		if err != nil {
			slog.Error("failed to init telemetry", "error", err)
			return
		}

		validatorx.InitValidator()
		err = config.LoadCustomValidations()
		if err != nil {
			slog.Error("failed to register custom validation", "error", err)
			return
		}

		infra.NewLog()

		serv, cleanUp := LoadServices()

		presentations.NewPresentation(serv, telemetry.Chain(cleanUp, otelCleanup))
	},
}

func main() {
	var rootCmd = &cobra.Command{}
	rootCmd.AddCommand(restApiCmd)
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
