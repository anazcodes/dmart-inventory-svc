package services

import (
	"context"
	"net/http"

	"github.com/anazibinurasheed/dmart-inventory-svc/internal/pb"
	"github.com/anazibinurasheed/dmart-inventory-svc/internal/usecase"
	"github.com/anazibinurasheed/dmart-inventory-svc/internal/usecase/interfaces"
	"github.com/anazibinurasheed/dmart-inventory-svc/internal/util"
)

type Server struct {
	InventoryUseCase interfaces.InventoryUseCase
	pb.UnimplementedInventoryServiceServer
}

func (s *Server) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CreateCategoryResponse, error) {
	err := s.InventoryUseCase.CreateCategory(ctx, req)
	if util.HasError(err) {
		return &pb.CreateCategoryResponse{
			Status: http.StatusNotAcceptable,
			Msg:    "failed",
			Error:  err.Error(),
		}, nil

	}

	return &pb.CreateCategoryResponse{
		Status: http.StatusCreated,
		Msg:    "success, created category " + req.Name,
	}, nil

}

func (s *Server) ReadCategories(ctx context.Context, req *pb.Request) (*pb.ReadCategoriesResponse, error) {
	datas, err := s.InventoryUseCase.ReadCategories(ctx, req)
	if util.HasError(err) {
		return &pb.ReadCategoriesResponse{
			Status:     http.StatusInternalServerError,
			Msg:        "failed, server error",
			Error:      err.Error(),
			Categories: datas,
		}, nil
	}

	return &pb.ReadCategoriesResponse{
		Status:     http.StatusOK,
		Msg:        "success",
		Categories: datas,
	}, nil
}

func (s *Server) AddProduct(ctx context.Context, req *pb.AddProductRequest) (*pb.AddProductResponse, error) {
	err := s.InventoryUseCase.AddProduct(ctx, req)
	if err == usecase.ErrRecordAlreadyExist {
		return &pb.AddProductResponse{
			Status: http.StatusConflict,
			Msg:    "failed to add product, a product already exist with the same name",
			Error:  err.Error(),
		}, nil
	}

	if util.HasError(err) {
		return &pb.AddProductResponse{
			Status: http.StatusNotAcceptable,
			Msg:    "failed to add product",
			Error:  err.Error(),
		}, nil
	}

	return &pb.AddProductResponse{
		Status: http.StatusCreated,
		Msg:    "success, product added",
	}, nil

}

func (s *Server) ReadProducts(ctx context.Context, req *pb.Request) (*pb.ReadProductsResponse, error) {
	data, err := s.InventoryUseCase.ReadProducts(ctx, req)
	if util.HasError(err) {
		return &pb.ReadProductsResponse{
			Status: http.StatusInternalServerError,
			Msg:    "failed",
			Error:  err.Error(),
		}, nil
	}

	return &pb.ReadProductsResponse{
		Status:   http.StatusOK,
		Msg:      "success",
		Products: data,
	}, nil
}
