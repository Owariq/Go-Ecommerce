package cartService

import (
	"github.com/Owariq/Go-Ecommerce/shopping-cart-service/internal/model"
	"github.com/Owariq/Go-Ecommerce/shopping-cart-service/internal/service"
	"github.com/Owariq/Go-Ecommerce/shopping-cart-service/internal/storage"
)

type serv struct {
	storage storage.CartStorage
}

func NewCartService(storage storage.CartStorage) service.CartService {
	return &serv{
		storage: storage,
	}
}

func (s *serv) CreateCart(userID int64) (int64, error) {
	return s.storage.CreateCart(userID)
}
func (s *serv) GetCart(cartID int64) (model.Cart, error) {
	return s.storage.GetCart(cartID)
}
func (s *serv) AddItem(cartID int64, item model.Item) (model.Cart, error) {
	return s.storage.AddItem(cartID, item)
}
func (s *serv) RemoveItem(cartID int64, productID int64) (model.Cart, error) {
	return s.storage.RemoveItem(cartID, productID)
}
func (s *serv) ClearCart(cartID int64) error {
	return s.storage.ClearCart(cartID)
}
