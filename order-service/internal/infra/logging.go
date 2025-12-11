package infra

import (
	"order-service/internal/config"

	"github.com/CROWNIX/go-utils/observability"
	"github.com/CROWNIX/go-utils/observability/loghook"
)

func NewLog() {
	observability.NewLog(observability.LogConfig{
		Mode:          "json",
		Level:         "debug",
		Env:           config.GetConfig().AppEnv,
		ServiceName:   config.GetConfig().AppName,
		ZerologStdOut: config.GetConfig().AppEnv == "development",
		SlogHook:      loghook.NewRotatingWriter("app.log", 10, 1, 3, true),
	})
}
