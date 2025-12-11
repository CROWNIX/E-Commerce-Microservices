package main

import (
	"log/slog"
	"order-service/internal/config"
	"order-service/internal/infra"
	"order-service/internal/presentations"

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
		validatorx.InitValidator()
		err := config.LoadCustomValidations()
		if err != nil {
			slog.Error("failed to register custom validation", "error", err)
			return
		}

		infra.NewLog()

		serv, cleanUp, err := LoadServices()
		if err != nil {
			panic(err)
		}

		presentations.NewPresentation(serv, cleanUp)
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
