package app

import (
	"log"
	"os"

	"github.com/Owariq/Go-Ecommerce/shopping-cart-service/internal/api"
	"github.com/Owariq/Go-Ecommerce/shopping-cart-service/internal/model"
	"github.com/Owariq/Go-Ecommerce/shopping-cart-service/internal/service"
	"github.com/Owariq/Go-Ecommerce/shopping-cart-service/internal/service/cartService"
	"github.com/Owariq/Go-Ecommerce/shopping-cart-service/internal/storage"
	"github.com/Owariq/Go-Ecommerce/shopping-cart-service/internal/storage/cartStorage"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type serviceProvider struct {
	dbClient    *gorm.DB
	cartRepo    storage.CartStorage
	cartService service.CartService

	cartImpl *api.CartApi
}

func NewServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) DBClient() *gorm.DB {
	dsn := os.Getenv("PG_DSN")

	if s.dbClient == nil {

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to database", err)
		}

		if err := db.AutoMigrate(&model.Cart{}); err != nil {
			log.Fatal("Failed to migrate database", err)
		}
		s.dbClient = db
	}
	return s.dbClient
}

func (s *serviceProvider) CartRepo() storage.CartStorage {
	if s.cartRepo == nil {
		s.cartRepo = cartStorage.NewCartStorage(s.DBClient())
	}
	return s.cartRepo
}

func (s *serviceProvider) CartService() service.CartService {
	if s.cartService == nil {
		s.cartService = cartService.NewCartService(s.CartRepo())
	}
	return s.cartService
}

func (s *serviceProvider) CartImpl() *api.CartApi {
	if s.cartImpl == nil {
		s.cartImpl = api.NewCartApi(s.CartService())
	}
	return s.cartImpl
}
