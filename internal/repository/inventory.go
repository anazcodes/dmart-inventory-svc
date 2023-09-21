package repository

import (
	"context"

	"github.com/anazibinurasheed/dmart-inventory-svc/internal/models"
	"github.com/anazibinurasheed/dmart-inventory-svc/internal/repository/interfaces"
	"github.com/anazibinurasheed/dmart-inventory-svc/internal/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Inventory struct {
	CollProduct  *mongo.Collection
	CollCategory *mongo.Collection
}

func NewInventoryRepo(db *mongo.Database) interfaces.InventoryRepo {
	return &Inventory{
		CollProduct: db.Collection("product"),

		CollCategory: db.Collection("category"),
	}
}

func (i *Inventory) CreateCategory(ctx context.Context, req models.Category) error {
	result, err := i.CollCategory.InsertOne(context.TODO(), req)
	if util.HasError(err) {
		return err
	}

	util.Logger(result.InsertedID)
	return nil
}

func (i *Inventory) ReadCategories(ctx context.Context) ([]models.Category, error) {
	var result []models.Category

	cursor, err := i.CollCategory.Find(context.TODO(), bson.M{})

	if util.HasError(err) {
		return []models.Category{}, err
	}

	err = cursor.All(context.TODO(), &result)
	if util.HasError(err) {
		return []models.Category{}, err
	}

	util.Logger(result)
	return result, nil
}

func (i *Inventory) DeleteCategoryByID(ctx context.Context, categoryID primitive.ObjectID) error {
	result, err := i.CollCategory.DeleteOne(context.TODO(), bson.M{"category_id": categoryID})

	if util.HasError(err) {
		return err
	}

	util.Logger(result)
	return nil
}

func (i *Inventory) AddProduct(ctx context.Context, req models.Product) error {
	result, err := i.CollProduct.InsertOne(context.TODO(), req)
	util.Logger(result)
	return nil
}

func (i *Inventory) FindCategory(ctx context.Context, req models.Category) (models.Category, error) {
	var data models.Category
	err := i.CollProduct.FindOne(context.TODO(), req).Decode(&data)

	util.Logger(result)
	return nil
}

func (i *Inventory) ReadProduct(ctx context.Context) ([]models.Product, error) {
	var result []models.Product

	cursor, err := i.CollProduct.Find(context.TODO(), bson.M{})

	if util.HasError(err) {
		return []models.Product{}, err
	}

	err = cursor.All(context.TODO(), &result)
	if util.HasError(err) {
		return []models.Product{}, err
	}

	util.Logger(result)
	return result, nil
}

func (i *Inventory) DeleteProduct(ctx context.Context, productID primitive.ObjectID) error {
	result, err := i.CollProduct.DeleteOne(context.TODO(), bson.M{"_id": productID})
	if util.HasError(err) {
		return err
	}

	util.Logger(result)
	return nil
}
