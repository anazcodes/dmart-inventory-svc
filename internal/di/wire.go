package di

import (
	"github.com/anazibinurasheed/dmart-inventory-svc/internal/config"
	"github.com/anazibinurasheed/dmart-inventory-svc/internal/db"
	"github.com/anazibinurasheed/dmart-inventory-svc/internal/repository"
	"github.com/anazibinurasheed/dmart-inventory-svc/internal/services"
	"github.com/anazibinurasheed/dmart-inventory-svc/internal/usecase"
)

func InitializeService(config config.Config) (*services.Server, error) {
	mongo, err := db.InitDB(config.MongoUrl)
	if err != nil {
		return &services.Server{}, err
	}

	inventoryRepo := repository.NewInventoryRepo(mongo)
	InventoryUseCase := usecase.NewInventoryUseCase(inventoryRepo)

	return &services.Server{
		InventoryUseCase: InventoryUseCase,
	}, nil

}
