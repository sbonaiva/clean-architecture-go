package usecase

import (
	"github.com/sbonaiva/clean-architecture-go/core/domain"
	"github.com/sbonaiva/clean-architecture-go/core/gateway"
	"github.com/sbonaiva/clean-architecture-go/util"
)

type produceCustomerEventUseCase struct {
	gateway gateway.ProduceCustomerEventGateway
	logger  util.Logger
}

type ProduceCustomerEventUseCase interface {
	Execute(event *domain.CustomerEvent) error
}

func NewProduceCustomerEventUseCase(g gateway.ProduceCustomerEventGateway, l util.Logger) ProduceCustomerEventUseCase {
	return &produceCustomerEventUseCase{
		gateway: g,
		logger:  l,
	}
}

func (u *produceCustomerEventUseCase) Execute(event *domain.CustomerEvent) error {
	u.logger.Info("producing customer event", "event", event.Type, "id", event.Customer.ID)
	return u.gateway.Execute(event)
}
