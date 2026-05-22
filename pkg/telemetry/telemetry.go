package telemetry

import (
	"context"

	"github.com/CROWNIX/go-utils/observability"
)

type Config struct {
	Enabled      bool
	ServiceName  string
	Env          string
	OtlpEndpoint string
	OtlpUser     string
	OtlpPassword string
}

func Init(ctx context.Context, cfg Config) (func(), error) {
	if !cfg.Enabled {
		return func() {}, nil
	}

	if cfg.OtlpEndpoint == "" {
		cfg.OtlpEndpoint = "localhost:4317"
	}

	return observability.NewObservabilityOtel(observability.OptionParams{
		ServiceName:  cfg.ServiceName,
		Env:          cfg.Env,
		OtlpEndpoint: cfg.OtlpEndpoint,
		OtlpUsername: cfg.OtlpUser,
		OtlpPassword: cfg.OtlpPassword,
	})
}

func Chain(cleanups ...func()) func() {
	return func() {
		for i := len(cleanups) - 1; i >= 0; i-- {
			if cleanups[i] != nil {
				cleanups[i]()
			}
		}
	}
}
