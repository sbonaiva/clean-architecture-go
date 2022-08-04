package logging

import "go.uber.org/zap"

func NewLoggerTool() *zap.SugaredLogger {

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	return logger.Sugar()
}
