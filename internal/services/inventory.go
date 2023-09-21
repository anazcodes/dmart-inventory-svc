package services

import (
	"context"
	"net/http"

	"github.com/anazibinurasheed/dmart-inventory-svc/internal/pb"
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
		util.Logger(err)
		return &pb.CreateCategoryResponse{
			Status: http.StatusNotAcceptable, //change it accordingly
			Msg:    "failed",
			Error:  err.Error(),
		}, nil

	}

	return &pb.CreateCategoryResponse{
		Status: http.StatusCreated,
		Msg:    "category created",
	}, nil

}
