package interfaces

import (
	"context"

	"github.com/anazibinurasheed/dmart-inventory-svc/internal/models"
	"github.com/anazibinurasheed/dmart-inventory-svc/internal/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InventoryRepo interface {
	// Category
	CreateCategory(ctx context.Context, req models.Category) error
	GetCategoryByName(ctx context.Context, name string) (models.Category, error)
	GetCategoryByID(ctx context.Context, categoryID string) (models.Category, error)
	ReadCategories(ctx context.Context, req models.PageInfo) ([]*pb.Category, error)
	DeleteCategoryByID(ctx context.Context, categoryID string) error

	// Product
	GetProductByName(ctx context.Context, name string) (models.Product, error)
	AddProduct(ctx context.Context, req models.Product) error
	ReadProducts(ctx context.Context, req models.PageInfo) ([]*pb.Product, error)
	DeleteProduct(ctx context.Context, productID primitive.ObjectID) error
}
