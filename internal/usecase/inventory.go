package usecase

import (
	"context"
	"errors"

	"github.com/anazibinurasheed/dmart-inventory-svc/internal/models"
	"github.com/anazibinurasheed/dmart-inventory-svc/internal/pb"

	"github.com/anazibinurasheed/dmart-inventory-svc/internal/repo"
	"github.com/anazibinurasheed/dmart-inventory-svc/internal/repo/interfaces"
	usecase "github.com/anazibinurasheed/dmart-inventory-svc/internal/usecase/interfaces"
	"github.com/anazibinurasheed/dmart-inventory-svc/internal/util"
)

var (
	ErrInputIsNotEnough   = errors.New("the request has not enough details")
	ErrRecordAlreadyExist = errors.New("record already exist")
	ErrCategoryNotExist   = errors.New("category does not exist")
)

type inventoryUseCase struct {
	Repo interfaces.InventoryRepo
}

func NewInventoryUseCase(inventoryRepo interfaces.InventoryRepo) usecase.InventoryUseCase {
	return &inventoryUseCase{
		Repo: inventoryRepo,
	}
}

func (i *inventoryUseCase) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) error {
	data, err := i.Repo.GetCategoryByName(ctx, req.Name)
	if err != nil {
		if err != repo.ErrNoDocuments {
			util.Logger(err)
			return err
		}
	}

	if data.Name != "" {
		return ErrRecordAlreadyExist
	}

	err = i.Repo.CreateCategory(ctx, models.Category{
		Name: req.Name,
	})
	if util.HasError(err) {
		return err
	}

	return nil
}

func (i *inventoryUseCase) ReadCategories(ctx context.Context, req *pb.Request) ([]*pb.Category, error) {
	skip, limit := util.Paginate(req.Page, req.Count)

	datas, err := i.Repo.ReadCategories(ctx, models.PageInfo{
		Skip:  skip,
		Limit: limit,
	})
	if util.HasError(err) {
		if err != repo.ErrNoDocuments {
			return nil, err
		}
	}

	return datas, nil
}

func (i *inventoryUseCase) AddProduct(ctx context.Context, req *pb.AddProductRequest) error {
	category, err := i.Repo.GetCategoryByID(ctx, req.CategoryID)
	util.Logger(category.ID, category.ID.Hex())

	if category.Name == "" {
		return ErrCategoryNotExist
	}

	if util.HasError(err) {
		return err
	}

	product, err := i.Repo.GetProductByName(ctx, req.Name)
	if util.HasError(err) {
		if err != repo.ErrNoDocuments {
			return err
		}
	}

	if product.Name != "" {
		return ErrRecordAlreadyExist
	}

	err = i.Repo.AddProduct(ctx, models.Product{
		CategoryID:  category.ID,
		Name:        req.Name,
		Brand:       category.Name,
		Description: req.Description,
		Image:       req.Image,
		Price:       req.Price,
	})

	return err
}

func (i *inventoryUseCase) ReadProducts(ctx context.Context, req *pb.Request) ([]*pb.Product, error) {
	skip, limit := util.Paginate(req.Page, req.Count)

	products, err := i.Repo.ReadProducts(ctx, models.PageInfo{
		Skip:  skip,
		Limit: limit,
	})
	if util.HasError(err) {
		if err != repo.ErrNoDocuments {
			return nil, err
		}
	}

	return products, nil

}
