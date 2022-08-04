package usecase

import (
	"github.com/sbonaiva/clean-architecture-go/core/domain"
	"github.com/sbonaiva/clean-architecture-go/core/gateway"
	"github.com/sbonaiva/clean-architecture-go/util"
)

type createCustomerUseCase struct {
	gateway      gateway.CreateCustomerGateway
	eventUseCase ProduceCustomerEventUseCase
	logger       util.Logger
}

type CreateCustomerUseCase interface {
	Execute(customer *domain.Customer) error
}

func NewCreateCustomerUseCase(g gateway.CreateCustomerGateway, u ProduceCustomerEventUseCase, l util.Logger) CreateCustomerUseCase {
	return &createCustomerUseCase{
		gateway:      g,
		eventUseCase: u,
		logger:       l,
	}
}

func (u *createCustomerUseCase) Execute(customer *domain.Customer) error {

	u.logger.Info("creating customer")

	errValidateCustomer := customer.Validate()

	if errValidateCustomer != nil {
		u.logger.Error("invalid customer fields", "error", errValidateCustomer.Error())
		return errValidateCustomer
	}

	errCreateCustomer := u.gateway.Execute(customer)

	if errCreateCustomer != nil {
		u.logger.Error("failed to create customer", "error", errCreateCustomer.Error())
		return errCreateCustomer
	}

	createCustomerEvent := domain.NewCreateCustomerEvent(customer)

	errProduceEvent := u.eventUseCase.Execute(createCustomerEvent)

	if errProduceEvent != nil {
		u.logger.Error("failed to produce create customer event", "error", errProduceEvent.Error())
	}

	return nil
}
