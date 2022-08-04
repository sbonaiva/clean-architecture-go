package gateway

import "github.com/sbonaiva/clean-architecture-go/core/domain"

type GetCustomerGateway interface {
	Execute(id string) (*domain.Customer, error)
}

type UpdateCustomerGateway interface {
	Execute(customer *domain.Customer) error
}

type DeleteCustomerGateway interface {
	Execute(id string) error
}

type CreateCustomerGateway interface {
	Execute(customer *domain.Customer) error
}

type ListCustomerGateway interface {
	Execute() ([]*domain.Customer, error)
}
