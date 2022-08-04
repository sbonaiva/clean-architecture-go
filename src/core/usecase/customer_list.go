package usecase

import (
	"github.com/sbonaiva/clean-architecture-go/core/domain"
	"github.com/sbonaiva/clean-architecture-go/core/gateway"
	"github.com/sbonaiva/clean-architecture-go/util"
)

type listCustomerUseCase struct {
	gateway gateway.ListCustomerGateway
	logger  util.Logger
}

type ListCustomerUseCase interface {
	Execute() ([]*domain.Customer, error)
}

func NewListCustomerUseCase(g gateway.ListCustomerGateway, l util.Logger) ListCustomerUseCase {
	return &listCustomerUseCase{
		gateway: g,
		logger:  l,
	}
}

func (u *listCustomerUseCase) Execute() ([]*domain.Customer, error) {
	u.logger.Info("listing customers")
	return u.gateway.Execute()
}
