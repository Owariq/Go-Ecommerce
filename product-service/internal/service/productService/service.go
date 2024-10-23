package productService

import (
	"context"

	"github.com/Owariq/Go-Ecommerce/product-service/internal/model"
	"github.com/Owariq/Go-Ecommerce/product-service/internal/service"
	"github.com/Owariq/Go-Ecommerce/product-service/internal/storage"
)

type serv struct {
	storage storage.ProductStorage
}

func NewProductService(storage storage.ProductStorage) service.ProductService {
	return &serv{storage: storage}
}


func (s *serv) Create(ctx context.Context, product model.Product) (int64, error) {

	return s.storage.Create(product)
}

func (s *serv) Get(ctx context.Context, id int64) (model.Product, error) {

	return s.storage.Get(ctx, id)
}

func (s *serv) Update(ctx context.Context, updatedProduct model.Product) (int64, error) {

	return s.storage.Update(updatedProduct)
}

func (s *serv) Delete(ctx context.Context, id int64) (int64, error) {

	return s.storage.Delete(ctx, id)
}

func (s *serv) List(ctx context.Context) ([]model.Product, error) {

	return s.storage.List()
}