package repo

import (
	"context"

	"github.com/anazibinurasheed/dmart-inventory-svc/internal/models"
	"github.com/anazibinurasheed/dmart-inventory-svc/internal/pb"
	"github.com/anazibinurasheed/dmart-inventory-svc/internal/repo/interfaces"
	"github.com/anazibinurasheed/dmart-inventory-svc/internal/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ErrNoDocuments = mongo.ErrNoDocuments
)

type inventory struct {
	Product  *mongo.Collection
	Category *mongo.Collection
}

func NewInventoryRepo(db *mongo.Database) interfaces.InventoryRepo {
	return &inventory{
		Product: db.Collection("product"),

		Category: db.Collection("category"),
	}
}

func (i *inventory) CreateCategory(ctx context.Context, req models.Category) error {
	result, err := i.Category.InsertOne(context.TODO(), req)
	util.Logger(result.InsertedID)

	return err
}

// this function will return ErrNoDocument if the category not found else the category
func (i *inventory) GetCategoryByName(ctx context.Context, name string) (models.Category, error) {
	var data models.Category
	filter := bson.M{"name": name}
	err := i.Category.FindOne(context.TODO(), filter).Decode(&data)
	return data, err
}

func (i *inventory) ReadCategories(ctx context.Context, req models.PageInfo) ([]*pb.Category, error) {

	options := options.Find().SetSkip(req.Skip).SetLimit(req.Limit)

	cursor, err := i.Category.Find(ctx, bson.M{}, options)

	if util.HasError(err) {
		return nil, err
	}
	defer cursor.Close(ctx)

	var datas []*pb.Category

	for cursor.Next(ctx) {
		var data models.Category
		err := cursor.Decode(&data)
		if util.HasError(err) {
			return nil, err
		}

		datas = append(datas, &pb.Category{
			ID:   data.ID.Hex(),
			Name: data.Name,
		})

		if util.HasError(cursor.Err()) {
			return nil, cursor.Err()
		}
	}

	return datas, nil
}

func (i *inventory) GetCategoryByID(ctx context.Context, categoryID string) (models.Category, error) {
	var data models.Category
	ID, err := primitive.ObjectIDFromHex(categoryID)

	if util.HasError(err) {
		return models.Category{}, err
	}
	util.Logger(categoryID, "<--req,obj--->", ID)
	filter := bson.M{"_id": ID}
	err = i.Category.FindOne(context.TODO(), filter).Decode(&data)

	return data, err
}

func (i *inventory) DeleteCategoryByID(ctx context.Context, categoryID string) error {

	ID, err := primitive.ObjectIDFromHex(categoryID)
	if util.HasError(err) {
		return err
	}

	result, err := i.Category.DeleteOne(context.TODO(),
		bson.M{"category_id": ID})

	util.Logger(result.DeletedCount)
	return err
}

func (i *inventory) GetProductByName(ctx context.Context, name string) (models.Product, error) {
	var data models.Product
	filter := bson.M{"name": name}
	err := i.Product.FindOne(context.TODO(), filter).Decode(&data)
	util.Logger(data, "--", name, "--", err)
	return data, err
}

func (i *inventory) AddProduct(ctx context.Context, req models.Product) error {
	result, err := i.Product.InsertOne(context.TODO(), req)
	util.Logger(result.InsertedID)
	return err
}

func (i *inventory) ReadProducts(ctx context.Context, req models.PageInfo) ([]*pb.Product, error) {
	options := options.Find().SetSkip(req.Skip).SetLimit(req.Limit)

	cursor, err := i.Product.Find(context.TODO(), bson.M{}, options)

	if util.HasError(err) {
		return nil, err
	}
	defer cursor.Close(ctx)

	var datas []*pb.Product

	for cursor.Next(context.TODO()) {
		var data models.Product
		err := cursor.Decode(&data)

		if util.HasError(err) {
			return nil, err
		}

		datas = append(datas, &pb.Product{
			Id:          data.ID.Hex(),
			CategoryID:  data.CategoryID.Hex(),
			Name:        data.Name,
			Brand:       data.Brand,
			Description: data.Description,
			Image:       data.Image,
			Price:       data.Price,
			IsBlocked:   data.IsBlocked,
		})

		if util.HasError(cursor.Err()) {
			return nil, cursor.Err()
		}

	}
	return datas, nil

}

func (i *inventory) DeleteProduct(ctx context.Context, productID primitive.ObjectID) error {
	result, err := i.Product.DeleteOne(context.TODO(), bson.M{"_id": productID})
	util.Logger(result)
	return err
}
