package storage

import (
	"context"

	"github.com/Owariq/Go-Ecommerce/product-service/internal/model"
)

type ProductStorage interface {
	Create(product model.Product) (int64, error)
	Get(ctx context.Context,id int64) ( model.Product, error)
	Update(updatedProduct  model.Product) (int64, error)
	Delete(ctx context.Context, id int64) (int64, error)
	List() ([]model.Product, error)
}