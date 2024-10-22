package service

import (
	"context"

	"github.com/Owariq/Go-Ecommerce/product-service/internal/model"
)

type ProductService interface {
	Create(ctx context.Context, product model.Product) (int64, error)
	Get(ctx context.Context, id int64) (model.Product, error)
	Update(ctx context.Context, updatedProduct model.Product) (int64, error)
	Delete(ctx context.Context, id int64) (int64, error)
	List(ctx context.Context) ([]model.Product, error)
}