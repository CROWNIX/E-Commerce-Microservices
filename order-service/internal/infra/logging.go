package infra

import (
	"order-service/internal/config"

	"github.com/CROWNIX/go-utils/observability"
)

func NewLog() {
	observability.NewLog(observability.LogConfig{
		Mode:          "json",
		Level:         "debug",
		Env:           config.GetConfig().AppEnv,
		ServiceName:   config.GetConfig().AppName,
		ZerologStdOut: config.GetConfig().AppEnv == "development",
	})
}
