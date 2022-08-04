package usecase

import (
	"github.com/sbonaiva/clean-architecture-go/core/domain"
	"github.com/sbonaiva/clean-architecture-go/core/gateway"
	"github.com/sbonaiva/clean-architecture-go/util"
)

type updateCustomerUseCase struct {
	gateway gateway.UpdateCustomerGateway
	logger  util.Logger
}

type UpdateCustomerUseCase interface {
	Execute(customer *domain.Customer) error
}

func NewUpdateCustomerUseCase(g gateway.UpdateCustomerGateway, l util.Logger) UpdateCustomerUseCase {
	return &updateCustomerUseCase{
		gateway: g,
		logger:  l,
	}
}

func (u *updateCustomerUseCase) Execute(customer *domain.Customer) error {
	u.logger.Info("updating customer", "id", customer.ID)
	return u.gateway.Execute(customer)
}
