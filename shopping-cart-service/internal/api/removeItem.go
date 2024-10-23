package api

import (
	"context"

	"github.com/Owariq/Go-Ecommerce/shopping-cart-service/internal/utils/converter"
	"github.com/Owariq/Go-Ecommerce/shopping-cart-service/pkg/cart_v1"
)

func (s *CartApi) RemoveItem(ctx context.Context, req *cart_v1.RemoveItemRequest) (*cart_v1.RemoveItemResponse, error) {

	cart, err := s.CartService.RemoveItem(req.GetCartId(), req.GetProductId())
	if err != nil {
		return nil, err
	}
	return &cart_v1.RemoveItemResponse{Cart: converter.FromModelCart(cart)}, nil
}
