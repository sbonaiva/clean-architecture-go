package repository

import (
	"context"
	"net/http"

	"github.com/sbonaiva/clean-architecture-go/core/domain"
	"github.com/sbonaiva/clean-architecture-go/core/gateway"
	"github.com/sbonaiva/clean-architecture-go/interface/repository/entity"
	"github.com/sbonaiva/clean-architecture-go/util"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CreateCustomerFaield = "Failed to create customer"
)

type createCustomerRepository struct {
	collection *mongo.Collection
	logger     util.Logger
}

type CreateCustomerRepository interface {
	gateway.CreateCustomerGateway
}

func NewCreateCustomerRepository(c *mongo.Collection, l util.Logger) CreateCustomerRepository {
	return &createCustomerRepository{
		collection: c,
		logger:     l,
	}
}

func (r *createCustomerRepository) Execute(customer *domain.Customer) error {

	entity := entity.NewCustomerEntity(customer)

	res, err := r.collection.InsertOne(context.Background(), entity)

	if err != nil {
		r.logger.Error("failed to create customer", "error", err.Error())
		return domain.NewCoreError(http.StatusInternalServerError, CreateCustomerFaield)
	}

	customer.ID = res.InsertedID.(primitive.ObjectID).Hex()

	return nil
}
