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
	GetCustomerInvalid  = "Invalid customer ID"
	GetCustomerNotFound = "Customer not found"
	GetCustomerFailed   = "Failed to get customer"
)

type getCustomerRepository struct {
	collection *mongo.Collection
	logger     util.Logger
}

type GetCustomerRepository interface {
	gateway.GetCustomerGateway
}

func NewGetCustomerRepository(c *mongo.Collection, l util.Logger) GetCustomerRepository {
	return &getCustomerRepository{
		collection: c,
		logger:     l,
	}
}

func (r *getCustomerRepository) Execute(id string) (*domain.Customer, error) {

	objectID, errHex := primitive.ObjectIDFromHex(id)

	if errHex != nil {
		r.logger.Error("invalid customer id", "error", errHex.Error())
		return nil, domain.NewCoreError(http.StatusBadRequest, GetCustomerInvalid)
	}

	var e entity.CustomerEntity

	errFind := r.collection.
		FindOne(context.Background(), bson.M{"_id": objectID}).
		Decode(&e)

	if errFind != nil {

		if errFind == mongo.ErrNoDocuments {
			r.logger.Error("customer not found", "error", errFind.Error())
			return nil, domain.NewCoreError(http.StatusNotFound, GetCustomerNotFound)
		}

		r.logger.Error("failed to retrieve customer", "error", errFind.Error())
		return nil, domain.NewCoreError(http.StatusNotFound, GetCustomerFailed)
	}

	return entity.FromCustomerEntity(e), nil
}
