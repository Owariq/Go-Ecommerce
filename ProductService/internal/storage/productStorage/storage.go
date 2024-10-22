package productStorage

import (
	"context"
	"fmt"
	"time"

	"github.com/Owariq/Go-Ecommerce/product-service/internal/model"
	"github.com/Owariq/Go-Ecommerce/product-service/internal/storage"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
	rdb *redis.Client
}

func NewProductStorage(db *gorm.DB, rdb *redis.Client) storage.ProductStorage {
	return &repo{db: db, rdb: rdb}
}

func (r *repo) Create(product model.Product) (int64, error) {
    
	err := r.db.Create(&product)
	if err.Error != nil {
		return 0, fmt.Errorf("failed to create product: %w", err.Error)
	}
	return product.ID, nil
}

func (r *repo) Get(ctx context.Context, id int64) (model.Product, error) {
	var product model.Product

   err := r.rdb.Get(ctx, fmt.Sprintf("product:%d", id)).Scan(&product)
    if err == nil {
		fmt.Printf("found product in redis: %v\n", product)
        return product, nil
    }
   
	err = r.db.First(&product, id).Error
	if err != nil {
		return model.Product{}, fmt.Errorf("failed to get product: %w", err)
	}
    err = r.rdb.Set(ctx, fmt.Sprintf("product:%d", id), product, 30 * time.Second).Err()
    if err != nil {
        fmt.Printf("failed to set product in redis: %v\n", err)
    }

	return product, nil
}

func (r *repo) Update(updatedProduct  model.Product) (int64, error) {
	var product model.Product

	r.db.First(&product, updatedProduct.ID)
	product = updatedProduct

	err := r.db.Save(&product)
	if err.Error != nil {
		return 0, fmt.Errorf("failed to update product: %w", err.Error)
	}
	return product.ID, nil
}

func (r *repo) Delete(ctx context.Context,id int64) (int64, error) {
	var product model.Product

	r.db.First(&product, id)

	err := r.db.Delete(&product).Error
	if err != nil {
		return 0, fmt.Errorf("failed to delete product: %w", err)
	}
	r.rdb.Del(ctx, fmt.Sprintf("product:%d", id))

	return product.ID, nil
}

func (r *repo) List() ([]model.Product, error) {
	var products []model.Product
	err := r.db.Find(&products)
	if err.Error != nil {
		return nil, fmt.Errorf("failed to list products: %w", err.Error)
	}
	return products, nil
}