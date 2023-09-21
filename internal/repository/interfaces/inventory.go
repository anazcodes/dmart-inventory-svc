package interfaces

import (
	"context"

	"github.com/anazibinurasheed/dmart-inventory-svc/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ctx = context.Context

type InventoryRepo interface {
	// Category
	CreateCategory(ctx, models.Category) error
	ReadCategories(ctx) ([]models.Category, error)
	DeleteCategoryByID(ctx, primitive.ObjectID) error

	// Product
	AddProduct(ctx, models.Product) error
	ReadProduct(ctx) ([]models.Product, error)
	DeleteProduct(ctx, primitive.ObjectID) error
}
