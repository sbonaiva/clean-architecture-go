package streaming

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sbonaiva/clean-architecture-go/infrastructure/configuration"
	"github.com/sbonaiva/clean-architecture-go/util"
)

func NewProducer(cfg configuration.Configuration, logger util.Logger) *kafka.Producer {

	producer, err := kafka.NewProducer(
		&kafka.ConfigMap{
			"bootstrap.servers": cfg.KAFKA_BROKERS,
			"client.id":         cfg.KAFKA_CLIENT_ID,
		},
	)

	if err != nil {
		logger.Panic("failed to create kafka producer", "error", err.Error())
	}

	defer producer.Close()

	return producer
}
