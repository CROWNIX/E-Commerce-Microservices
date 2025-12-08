package workers

import (
	"context"

	"github.com/CROWNIX/E-Commerce-Microservices/api-gateway/internal/config"
	"github.com/CROWNIX/E-Commerce-Microservices/api-gateway/internal/server"
)

func runHttpWorker(cfg *config.Config, ctx context.Context) {
	srv := server.NewHttpServer(cfg)
	go srv.Start()

	<-ctx.Done()
	srv.Shutdown()
}
