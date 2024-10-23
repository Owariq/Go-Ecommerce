package api

import (
	"github.com/Owariq/Go-Ecommerce/shopping-cart-service/internal/service"
	"github.com/Owariq/Go-Ecommerce/shopping-cart-service/pkg/cart_v1"
)

type CartApi struct {
	cart_v1.UnimplementedProductV1Server
	CartService service.CartService
}

func NewCartApi(cartService service.CartService) *CartApi {
	return &CartApi{
		CartService: cartService,
	}
}
