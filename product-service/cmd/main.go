package main

import (
	"context"
	"log/slog"
	"net"
	"product-service/internal/config"
	"product-service/internal/infra"
	"product-service/internal/presentations"
	grpcPresentation "product-service/internal/presentations/grpc"
	pb "pkg/proto/generated/product"
	"pkg/telemetry"

	"github.com/CROWNIX/go-utils/validatorx"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
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

		// Start gRPC server
		go func() {
			lis, err := net.Listen("tcp", ":50051")
			if err != nil {
				slog.Error("failed to listen", "error", err)
				return
			}
			s := grpc.NewServer(telemetry.GRPCServerOptions()...)
			pb.RegisterProductServiceServer(s, grpcPresentation.NewServer(serv.ProductService))
			slog.Info("gRPC server listening at :50051")
			if err := s.Serve(lis); err != nil {
				slog.Error("failed to serve", "error", err)
			}
		}()

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
