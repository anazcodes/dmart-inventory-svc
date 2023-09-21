package usecase

import (
	"context"
	"fmt"

	"github.com/anazibinurasheed/dmart-inventory-svc/internal/models"
	"github.com/anazibinurasheed/dmart-inventory-svc/internal/pb"
	"github.com/anazibinurasheed/dmart-inventory-svc/internal/repository/interfaces"
	usecase "github.com/anazibinurasheed/dmart-inventory-svc/internal/usecase/interfaces"
)

type InventoryUseCase struct {
	UseCase interfaces.InventoryRepo
}

func NewInventoryUseCase(inventoryRepo interfaces.InventoryRepo) usecase.InventoryUseCase {
	return &InventoryUseCase{
		UseCase: inventoryRepo,
	}
}

func (i *InventoryUseCase) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) error {

	fmt.Println(i.UseCase.CreateCategory(ctx, models.Category{
		Name: req.Name,
	}))
	return nil
}
