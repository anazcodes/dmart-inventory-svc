package interfaces

import (
	"context"

	"github.com/anazibinurasheed/dmart-inventory-svc/internal/pb"
)

type InventoryUseCase interface {
	CreateCategory(context.Context, *pb.CreateCategoryRequest) error
	ReadCategories(ctx context.Context, req *pb.Request) ([]*pb.Category, error)
	AddProduct(ctx context.Context, req *pb.AddProductRequest) error
	ReadProducts(ctx context.Context, req *pb.Request) ([]*pb.Product, error)
}
