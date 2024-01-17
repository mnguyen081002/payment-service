package repository

import (
	"paymentservice/config"
	"paymentservice/internal/domain"
	"paymentservice/internal/infrastructure"
	"paymentservice/internal/repository/gormlib"
	"paymentservice/internal/repository/mongo"

	"go.uber.org/fx"
)

type UnitOfWork struct {
	ProductAttributesRepository domain.ProductAttributesRepository
	TierVariationRepository     domain.TierVariationRepository
	ProductModelRepository      domain.ProductModelRepository
	ProductRepository           domain.ProductRepository
	CategoryRepository          domain.CategoryRepository
	RatingRepository            domain.RatingRepository
}

func NewUnitOfWorkGorm() *UnitOfWork {
	return &UnitOfWork{
		ProductRepository:           gormlib.NewProductRepository(),
		ProductAttributesRepository: gormlib.NewProductAttributesRepository(),
		TierVariationRepository:     gormlib.NewTierVariationRepository(),
		ProductModelRepository:      gormlib.NewProductModelRepository(),
		CategoryRepository:          gormlib.NewCategoryRepository(),
		RatingRepository:            gormlib.NewRatingRepository(),
	}
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
