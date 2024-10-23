package service

import "github.com/Owariq/Go-Ecommerce/shopping-cart-service/internal/model"

type CartService interface {
	CreateCart(userID int64) (int64, error)
	GetCart(cartID int64) (model.Cart, error)
	AddItem(cartID int64, item model.Item) (model.Cart, error)
	RemoveItem(cartID int64, productID int64) (model.Cart, error)
	ClearCart(cartID int64) error
}
