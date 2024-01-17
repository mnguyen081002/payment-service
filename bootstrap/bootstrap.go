package bootstrap

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"paymentservice/config"
	controller "paymentservice/internal/api/controllers"
	"paymentservice/internal/api/middlewares"
	"paymentservice/internal/api/route"
	"paymentservice/internal/infrastructure"
	"paymentservice/internal/lib"
	"paymentservice/internal/messaging"
	"paymentservice/internal/repository"
	service "paymentservice/internal/services"
	"paymentservice/internal/utils"
)

func inject() fx.Option {
	return fx.Options(
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
		//fx.NopLogger,
		fx.Provide(
			config.NewConfig("/config/config.yaml"),
			utils.NewTimeoutContext,
		),
		route.Module,
		lib.Module,
		repository.Module,
		service.Module,
		controller.Module,
		middlewares.Module,
		infrastructure.Module,
		messaging.Module,
	)
}

func Run() {
	fx.New(inject()).Run()
}
