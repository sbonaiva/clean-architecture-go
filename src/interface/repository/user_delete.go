package repository

import (
	"context"
	"net/http"

	"github.com/sbonaiva/clean-architecture-go/core/domain"
	"github.com/sbonaiva/clean-architecture-go/core/gateway"
	"github.com/sbonaiva/clean-architecture-go/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DeleteCustomerInvalid  = "Invalid coustumer ID"
	DeleteCustomerNotFound = "Customer not found"
	DeleteCustomerFailed   = "Failed to delete customer"
)

type deleteCustomerRepository struct {
	collection *mongo.Collection
	logger     util.Logger
}

type DeleteCustomerRepository interface {
	gateway.DeleteCustomerGateway
}

func NewDeleteCustomerRepository(c *mongo.Collection, l util.Logger) DeleteCustomerRepository {
	return &deleteCustomerRepository{
		collection: c,
		logger:     l,
	}
}

func (r *deleteCustomerRepository) Execute(id string) error {

	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		r.logger.Error("invalid customer id", "error", err.Error())
		return domain.NewCoreError(http.StatusBadRequest, DeleteCustomerInvalid)
	}

	res, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": objectID})

	if err != nil {
		r.logger.Error("failed to delete customer", "error", err.Error())
		return domain.NewCoreError(http.StatusBadRequest, DeleteCustomerFailed)
	}

	if res.DeletedCount == 0 {
		r.logger.Error("customer not found", "error", err.Error())
		return domain.NewCoreError(http.StatusNotFound, DeleteCustomerNotFound)
	}

	return nil
}
