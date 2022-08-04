package database

import (
	"context"

	"github.com/sbonaiva/clean-architecture-go/infrastructure/configuration"
	"github.com/sbonaiva/clean-architecture-go/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(cfg configuration.Configuration, logger util.Logger) *mongo.Client {

	ctx := context.Background()

	opt := options.Client().ApplyURI(cfg.MONGO_URL)

	cli, err := mongo.Connect(ctx, opt)

	if err != nil {
		logger.Panic("failed to connect with mongo", "error", err.Error())
	}

	err = cli.Ping(ctx, nil)

	if err != nil {
		logger.Panic("failed to ping mongo", "error", err.Error())
	}

	return cli
}
