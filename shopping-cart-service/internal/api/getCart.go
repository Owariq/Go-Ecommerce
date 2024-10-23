package api

import (
	"context"

	"github.com/Owariq/Go-Ecommerce/shopping-cart-service/internal/utils/converter"
	"github.com/Owariq/Go-Ecommerce/shopping-cart-service/pkg/cart_v1"
)

func (s *CartApi) GetCart(ctx context.Context, req *cart_v1.GetCartRequest) (*cart_v1.GetCartResponse, error) {

	cart, err := s.CartService.GetCart(req.GetCartId())
	if err != nil {
		return nil, err
	}

	return &cart_v1.GetCartResponse{Cart: converter.FromModelCart(cart)}, nil
}
