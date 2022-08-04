package registry

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	"github.com/sbonaiva/clean-architecture-go/util"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	Database = "app"
)

type registry struct {
	client   *mongo.Client
	producer *kafka.Producer
	logger   util.Logger
}

type Registry interface {
	RegistryEndpoints(e *gin.Engine)
}

func NewRegistry(c *mongo.Client, p *kafka.Producer, l util.Logger) Registry {
	return &registry{
		client:   c,
		producer: p,
		logger:   l,
	}
}

func (r *registry) RegistryEndpoints(e *gin.Engine) {

	customerRegistry := NewCustomerRegistry(r.client.Database(Database), r.producer, r.logger)
	e.GET("/api/v1/customers/:id", customerRegistry.GetController().Handle)
	e.GET("/api/v1/customers", customerRegistry.ListController().Handle)
	e.POST("/api/v1/customers", customerRegistry.CreateController().Handle)
	e.PUT("/api/v1/customers/:id", customerRegistry.UpdateController().Handle)
	e.DELETE("/api/v1/customers/:id", customerRegistry.DeleteController().Handle)
}
