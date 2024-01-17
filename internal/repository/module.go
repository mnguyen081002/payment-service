package repository

import (
	"paymentservice/config"
	"paymentservice/internal/infrastructure"
	"paymentservice/internal/repository/gormlib"
	"paymentservice/internal/repository/mongo"

	"go.uber.org/fx"
)

type UnitOfWork struct {
}

func NewUnitOfWorkGorm() *UnitOfWork {
	return &UnitOfWork{}
}

func NewUnitOfWorkMongo() *UnitOfWork {
	return &UnitOfWork{}
}

func NewUnitOfWork(config *config.Config) *UnitOfWork {
	if config.Database.Driver == "mongo" {
		return NewUnitOfWorkMongo()
	} else {
		return NewUnitOfWorkGorm()
	}
}

func NewRepository(config *config.Config, database infrastructure.Database) infrastructure.DatabaseTransaction {
	if config.Database.Driver == "mongo" {
		return mongo.NewMongoTransaction(database)
	} else {
		return gormlib.NewGormTransaction(database)
	}
}

var Module = fx.Options(
	fx.Provide(NewUnitOfWork),
	fx.Provide(NewRepository),
)
