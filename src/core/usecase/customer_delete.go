package usecase

import (
	"github.com/sbonaiva/clean-architecture-go/core/gateway"
	"github.com/sbonaiva/clean-architecture-go/util"
)

type deleteCustomerUseCase struct {
	gateway gateway.DeleteCustomerGateway
	logger  util.Logger
}

type DeleteCustomerUseCase interface {
	Execute(id string) error
}

func NewDeleteCustomerUseCase(g gateway.DeleteCustomerGateway, l util.Logger) DeleteCustomerUseCase {
	return &deleteCustomerUseCase{
		gateway: g,
		logger:  l,
	}
}

func (u *deleteCustomerUseCase) Execute(id string) error {
	u.logger.Info("deleting customer", "id", id)
	return u.gateway.Execute(id)
}
