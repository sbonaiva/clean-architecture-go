package domain

const (
	CreateCustomer = "CreateCustomer"
)

type CustomerEvent struct {
	Type     string
	Customer *Customer
}

func NewCreateCustomerEvent(c *Customer) *CustomerEvent {
	return &CustomerEvent{
		Type:     CreateCustomer,
		Customer: c,
	}
}
