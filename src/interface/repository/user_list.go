package repository

import (
	"context"
	"net/http"

	"github.com/sbonaiva/clean-architecture-go/core/domain"
	"github.com/sbonaiva/clean-architecture-go/core/gateway"
	"github.com/sbonaiva/clean-architecture-go/interface/repository/entity"
	"github.com/sbonaiva/clean-architecture-go/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ListCustomerFailed = "Failed to list customers"
)

type listCustomerRepository struct {
	collection *mongo.Collection
	logger     util.Logger
}

type ListCustomerRepository interface {
	gateway.ListCustomerGateway
}

func NewListCustomerRepository(c *mongo.Collection, l util.Logger) ListCustomerRepository {
	return &listCustomerRepository{
		collection: c,
		logger:     l,
	}
}

func (r *listCustomerRepository) Execute() ([]*domain.Customer, error) {

	cur, err := r.collection.Find(context.Background(), bson.D{})

	if err != nil {
		r.logger.Error("failed to list customers", "error", err.Error())
		return nil, domain.NewCoreError(http.StatusInternalServerError, ListCustomerFailed)
	}

	var customers []entity.CustomerEntity = make([]entity.CustomerEntity, 0)

	for cur.Next(context.Background()) {

		var customer entity.CustomerEntity

		err := cur.Decode(&customer)

		if err != nil {
			r.logger.Warn("failed to decode customer", "warn", err.Error())
		} else {
			customers = append(customers, customer)
		}
	}

	return entity.FromCustomerEntitySlice(customers), nil
}
