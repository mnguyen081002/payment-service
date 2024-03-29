package middlewares

import (
	config "paymentservice/config"
	"paymentservice/internal/infrastructure"

	"go.uber.org/zap"
)

type GinMiddleware struct {
	logger *zap.Logger
	config *config.Config
	db     infrastructure.Database
}

func NewMiddleware(config *config.Config, db infrastructure.Database, logger *zap.Logger) *GinMiddleware {
	return &GinMiddleware{
		logger: logger,
		config: config,
		db:     db,
	}
}
