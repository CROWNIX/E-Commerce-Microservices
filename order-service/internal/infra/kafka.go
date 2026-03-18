package infra

import (
	"log/slog"
	"order-service/internal/config"

	utilKafka "github.com/CROWNIX/go-utils/broker/kafka"
)

func NewKafka() (utilKafka.PubSub, func(), error) {
    cfg := config.GetConfig()
    kafka := utilKafka.New(
        utilKafka.Writer(
            cfg.ServiceName,
            []string{cfg.Kafka.Broker},
        ),
    )

    slog.Info("kafka producer initialized")

    cleanup := func() {
        slog.Info("closing kafka producer")
        kafka.Close()
    }

    return kafka, cleanup, nil
}
