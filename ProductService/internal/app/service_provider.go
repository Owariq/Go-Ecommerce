package app

import (
	"log"
	"os"

	"github.com/Owariq/Go-Ecommerce/product-service/internal/api"
	"github.com/Owariq/Go-Ecommerce/product-service/internal/model"
	"github.com/Owariq/Go-Ecommerce/product-service/internal/service"
	"github.com/Owariq/Go-Ecommerce/product-service/internal/service/productService"
	"github.com/Owariq/Go-Ecommerce/product-service/internal/storage"
	"github.com/Owariq/Go-Ecommerce/product-service/internal/storage/productStorage"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type serviceProvider struct {

	dbClient    *gorm.DB
	rdbClient   *redis.Client
	productRepo    storage.ProductStorage
	productService service.ProductService

	productImpl *api.ProductApi
}

func NewServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (a *serviceProvider) DBClient() *gorm.DB {
dsn := os.Getenv("PG_DSN")

if a.dbClient == nil {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	if err := db.AutoMigrate(&model.Product{}); err != nil {
		log.Fatal("Failed to migrate database", err)
	}
	a.dbClient = db
	}
	return a.dbClient
}

func (a *serviceProvider) RDBClient() *redis.Client {
	if a.rdbClient == nil {
		rdb := redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_ADDR"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
		})
		a.rdbClient = rdb
	}
	return a.rdbClient
}

func (a *serviceProvider) ProductRepo() storage.ProductStorage {
	if a.productRepo == nil {
		a.productRepo = productStorage.NewProductStorage(a.DBClient(), a.RDBClient())
	}
	return a.productRepo
}

func (a *serviceProvider) ProductService() service.ProductService {
	if a.productService == nil {
		a.productService = productService.NewProductService(a.ProductRepo())
	}
	return a.productService
}

func (a *serviceProvider) ProductImpl() *api.ProductApi {
	if a.productImpl == nil {
		a.productImpl = api.NewProductApi(a.ProductService())
	}
	return a.productImpl
}
