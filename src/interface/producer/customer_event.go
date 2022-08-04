package producer

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sbonaiva/clean-architecture-go/core/domain"
	"github.com/sbonaiva/clean-architecture-go/core/gateway"
	"github.com/sbonaiva/clean-architecture-go/util"
)

var (
	customerTopic = "sbonaiva.customers.json"
)

type customerEventProducer struct {
	producer *kafka.Producer
	logger   util.Logger
}

type CustomerEventProducer interface {
	gateway.ProduceCustomerEventGateway
}

func NewCustomerEventProducer(p *kafka.Producer, l util.Logger) CustomerEventProducer {
	return &customerEventProducer{
		producer: p,
		logger:   l,
	}
}

func (p *customerEventProducer) Execute(event *domain.CustomerEvent) error {

	message, errMarshal := json.Marshal(event)

	if errMarshal != nil {
		p.logger.Error("failed to marshal event", "error", errMarshal.Error())
		return errMarshal
	}

	errProducer := p.producer.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &customerTopic,
				Partition: kafka.PartitionAny,
			},
			Value: message,
		}, nil,
	)

	if errProducer != nil {
		p.logger.Error("failed to produce event", "error", errProducer.Error())
		return errProducer
	}

	return nil
}
