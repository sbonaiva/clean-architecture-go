package gateway

import "github.com/sbonaiva/clean-architecture-go/core/domain"

type ProduceCustomerEventGateway interface {
	Execute(event *domain.CustomerEvent) error
}
