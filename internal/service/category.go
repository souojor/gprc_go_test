package service

import (
	"context"

	"github.com/devfullcycle/14-gRPC/internal/database"
	"github.com/devfullcycle/14-gRPC/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	categoty, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	categotyResponse := &pb.Category{
		Id:          categoty.ID,
		Name:        categoty.Name,
		Description: categoty.Description,
	}

	return categotyResponse, nil
}
