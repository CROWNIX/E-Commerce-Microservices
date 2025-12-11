package presentations

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"order-service/internal/config"
	"order-service/internal/presentations/handler"
	"order-service/internal/services"
	"os/signal"
	"syscall"
	"time"

	"github.com/CROWNIX/go-utils/http/server/ginx"
	"github.com/gin-gonic/gin"
)

func NewPresentation(service *services.Service, cleanUp func()) {
	gin.SetMode(config.GetConfig().GinMode)

	r := ginx.NewGin(ginx.GinConfig{
		UseOtel: false,
		AppName: config.GetConfig().AppName,
	})

	h := handler.NewHandler(handler.Options{
		Service: service,
	})

	h.RegisterRoutes(r)

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", config.GetConfig().RestApiPort),
		Handler:           r.Handler(),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	sigCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	<-sigCtx.Done()

	slog.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}

	cleanUp()
	slog.Info("Server exiting")
}
