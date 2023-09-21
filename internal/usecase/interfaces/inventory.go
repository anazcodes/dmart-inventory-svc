package interfaces

import (
	"context"

	"github.com/anazibinurasheed/dmart-inventory-svc/internal/pb"
)

type InventoryUseCase interface {
	CreateCategory(context.Context, *pb.CreateCategoryRequest) error
}
