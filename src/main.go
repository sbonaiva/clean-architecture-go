package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sbonaiva/clean-architecture-go/infrastructure/configuration"
	"github.com/sbonaiva/clean-architecture-go/infrastructure/database"
	"github.com/sbonaiva/clean-architecture-go/infrastructure/logging"
	"github.com/sbonaiva/clean-architecture-go/infrastructure/streaming"
	"github.com/sbonaiva/clean-architecture-go/registry"
	"github.com/sbonaiva/clean-architecture-go/util"
)

func main() {

	router := gin.Default()

	loggerTool := logging.NewLoggerTool()

	logger := util.NewLogger(loggerTool)

	configuration := configuration.NewConfiguration(logger)

	database := database.Connect(configuration, logger)

	producer := streaming.NewProducer(configuration, logger)

	registry := registry.NewRegistry(database, producer, logger)

	registry.RegistryEndpoints(router)

	router.Run(":8080")
}
