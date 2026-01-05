package config

import (
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewKafkaConsumer(config *viper.Viper, log *logrus.Logger) *kafka.Reader {
	brokers := []string{
		config.GetString("kafka.bootstrap.servers"),
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		GroupID: config.GetString("kafka.group.id"),
		Topic:   config.GetString("kafka.topic"),

		StartOffset: kafka.FirstOffset,
	})

	log.Info("Kafka consumer initialized")

	return reader
}

func NewKafkaProducer(config *viper.Viper, log *logrus.Logger) *kafka.Writer {
	brokers := []string{
		config.GetString("kafka.bootstrap.servers"),
	}

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  brokers,
		Balancer: &kafka.LeastBytes{},
	})

	log.Info("Kafka producer initialized")

	return writer
}
