package dto

import "github.com/sbonaiva/clean-architecture-go/core/domain"

type CustomerRequest struct {
	Email  string `json:"email"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

type CustomerResponse struct {
	ID     string `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

func FromCustomerRequest(c CustomerRequest) *domain.Customer {
	return &domain.Customer{
		Email:  c.Email,
		Name:   c.Name,
		Active: c.Active,
	}
}

func ToCustomerResponse(c *domain.Customer) *CustomerResponse {
	return &CustomerResponse{
		ID:     c.ID,
		Email:  c.Email,
		Name:   c.Name,
		Active: c.Active,
	}
}

func ToCustomerResponseSlice(c []*domain.Customer) []*CustomerResponse {

	var customers []*CustomerResponse = make([]*CustomerResponse, 0)

	for _, customer := range c {
		customers = append(customers, ToCustomerResponse(customer))
	}

	return customers
}
