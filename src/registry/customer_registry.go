package registry

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sbonaiva/clean-architecture-go/core/usecase"
	"github.com/sbonaiva/clean-architecture-go/interface/controller"
	"github.com/sbonaiva/clean-architecture-go/interface/producer"
	"github.com/sbonaiva/clean-architecture-go/interface/repository"
	"github.com/sbonaiva/clean-architecture-go/util"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CustomersCollection = "customers"
)

type customerRegistry struct {
	collection *mongo.Collection
	producer   *kafka.Producer
	logger     util.Logger
}

type CustomerRegistry interface {
	GetController() controller.GetCustomerController
	ListController() controller.ListCustomerController
	CreateController() controller.CreateCustomerController
	UpdateController() controller.UpdateCustomerController
	DeleteController() controller.DeleteCustomerController
}

func NewCustomerRegistry(d *mongo.Database, p *kafka.Producer, l util.Logger) CustomerRegistry {
	return &customerRegistry{
		collection: d.Collection(CustomersCollection),
		producer:   p,
		logger:     l,
	}
}

func (r *customerRegistry) GetController() controller.GetCustomerController {

	getCustomerRepository := repository.NewGetCustomerRepository(r.collection, r.logger)
	getCustomerUseCase := usecase.NewGetCustomerUseCase(getCustomerRepository, r.logger)
	return controller.NewGetCustomerController(getCustomerUseCase, r.logger)
}

func (r *customerRegistry) ListController() controller.ListCustomerController {

	listCustomerRepository := repository.NewListCustomerRepository(r.collection, r.logger)
	listCustomerUseCase := usecase.NewListCustomerUseCase(listCustomerRepository, r.logger)
	return controller.NewListCustomerController(listCustomerUseCase, r.logger)
}

func (r *customerRegistry) CreateController() controller.CreateCustomerController {

	eventProducer := producer.NewCustomerEventProducer(r.producer, r.logger)
	produceEventUseCase := usecase.NewProduceCustomerEventUseCase(eventProducer, r.logger)
	createCustomerRepository := repository.NewCreateCustomerRepository(r.collection, r.logger)
	createCustomerUseCase := usecase.NewCreateCustomerUseCase(createCustomerRepository, produceEventUseCase, r.logger)
	return controller.NewCreateCustomerController(createCustomerUseCase, r.logger)
}

func (r *customerRegistry) UpdateController() controller.UpdateCustomerController {

	updateCustomerRepository := repository.NewUpdateCustomerRepository(r.collection, r.logger)
	updateCustomerUseCase := usecase.NewUpdateCustomerUseCase(updateCustomerRepository, r.logger)
	return controller.NewUpdateCustomerController(updateCustomerUseCase, r.logger)
}

func (r *customerRegistry) DeleteController() controller.DeleteCustomerController {

	deleteCustomerRepository := repository.NewDeleteCustomerRepository(r.collection, r.logger)
	deleteCustomerUseCase := usecase.NewDeleteCustomerUseCase(deleteCustomerRepository, r.logger)
	return controller.NewDeleteCustomerController(deleteCustomerUseCase, r.logger)
}
