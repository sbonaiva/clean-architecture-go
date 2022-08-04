package configuration

import (
	"github.com/sbonaiva/clean-architecture-go/util"
	"github.com/spf13/viper"
)

type Configuration struct {
	MONGO_URL       string
	KAFKA_BROKERS   string
	KAFKA_CLIENT_ID string
}

func NewConfiguration(logger util.Logger) Configuration {

	viper.SetConfigName("local")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("env")

	errReadInConfig := viper.ReadInConfig()

	if errReadInConfig != nil {
		logger.Panic("configuration read in failed", "error", errReadInConfig.Error())
	}

	var configuration Configuration

	errUnmarshal := viper.Unmarshal(&configuration)

	if errUnmarshal != nil {
		logger.Panic("configuration unmarshal failed", "error", errUnmarshal.Error())
	}

	return configuration
}
