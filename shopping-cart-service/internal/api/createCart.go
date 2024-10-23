package api

import (
	"context"

	"github.com/Owariq/Go-Ecommerce/shopping-cart-service/pkg/cart_v1"
)

func (s *CartApi) CreateCart(ctx context.Context, req *cart_v1.CreateCartRequest) (*cart_v1.CreateCartResponse, error) {
	cart, err := s.CartService.CreateCart(req.GetUserId())
	if err != nil {
		return nil, err
	}
	return &cart_v1.CreateCartResponse{CartId: cart}, nil
}
