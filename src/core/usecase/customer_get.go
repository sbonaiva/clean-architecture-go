package usecase

import (
	"github.com/sbonaiva/clean-architecture-go/core/domain"
	"github.com/sbonaiva/clean-architecture-go/core/gateway"
	"github.com/sbonaiva/clean-architecture-go/util"
)

type getCustomerUseCase struct {
	gateway gateway.GetCustomerGateway
	logger  util.Logger
}

type GetCustomerUseCase interface {
	Execute(id string) (*domain.Customer, error)
}

func NewGetCustomerUseCase(g gateway.GetCustomerGateway, l util.Logger) GetCustomerUseCase {
	return &getCustomerUseCase{
		gateway: g,
		logger:  l,
	}
}

func (u *getCustomerUseCase) Execute(id string) (*domain.Customer, error) {
	u.logger.Info("retrieving customer", "id", id)
	return u.gateway.Execute(id)
}
