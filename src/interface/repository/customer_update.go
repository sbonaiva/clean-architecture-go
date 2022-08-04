package repository

import (
	"context"
	"net/http"

	"github.com/sbonaiva/clean-architecture-go/core/domain"
	"github.com/sbonaiva/clean-architecture-go/core/gateway"
	"github.com/sbonaiva/clean-architecture-go/interface/repository/entity"
	"github.com/sbonaiva/clean-architecture-go/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	UpdateCustomerInvalid  = "Invalid customer ID"
	UpdateCustomerNotFound = "Customer not found"
	UpdateCustomerFailed   = "Failed to update customer"
)

type updateCustomerRepository struct {
	collection *mongo.Collection
	logger     util.Logger
}

type UpdateCustomerRepository interface {
	gateway.UpdateCustomerGateway
}

func NewUpdateCustomerRepository(c *mongo.Collection, l util.Logger) UpdateCustomerRepository {
	return &updateCustomerRepository{
		collection: c,
		logger:     l,
	}
}

func (r *updateCustomerRepository) Execute(customer *domain.Customer) error {

	objectID, err := primitive.ObjectIDFromHex(customer.ID)

	if err != nil {
		r.logger.Error("invalid customer", "error", err.Error())
		return domain.NewCoreError(http.StatusBadRequest, UpdateCustomerInvalid)
	}

	entity := entity.ToCustomerEntity(objectID, customer)

	res, err := r.collection.ReplaceOne(context.Background(), bson.M{"_id": objectID}, entity)

	if err != nil {
		r.logger.Error("failed to update user", "error", err.Error())
		return domain.NewCoreError(http.StatusInternalServerError, UpdateCustomerFailed)
	}

	if res.MatchedCount == 0 {
		r.logger.Error("customer not found", "error", err.Error())
		return domain.NewCoreError(http.StatusNotFound, UpdateCustomerNotFound)
	}

	return nil
}
