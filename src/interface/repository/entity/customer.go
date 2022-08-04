package entity

import (
	"github.com/sbonaiva/clean-architecture-go/core/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomerEntity struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Email  string             `bson:"email"`
	Name   string             `bson:"name"`
	Active bool               `bson:"active"`
}

func NewCustomerEntity(d *domain.Customer) *CustomerEntity {

	return &CustomerEntity{
		Email:  d.Email,
		Name:   d.Name,
		Active: d.Active,
	}
}

func ToCustomerEntity(id primitive.ObjectID, d *domain.Customer) *CustomerEntity {

	return &CustomerEntity{
		ID:     id,
		Email:  d.Email,
		Name:   d.Name,
		Active: d.Active,
	}
}

func FromCustomerEntity(e CustomerEntity) *domain.Customer {
	return &domain.Customer{
		ID:     e.ID.Hex(),
		Email:  e.Email,
		Name:   e.Name,
		Active: e.Active,
	}
}

func FromCustomerEntitySlice(s []CustomerEntity) []*domain.Customer {

	var customers []*domain.Customer = make([]*domain.Customer, 0)

	for _, customer := range s {
		customers = append(customers, FromCustomerEntity(customer))
	}

	return customers
}
