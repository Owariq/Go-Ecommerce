package api

import (
	"context"
	"fmt"

	"github.com/Owariq/Go-Ecommerce/shopping-cart-service/internal/model"
	"github.com/Owariq/Go-Ecommerce/shopping-cart-service/internal/utils/converter"
	"github.com/Owariq/Go-Ecommerce/shopping-cart-service/pkg/cart_v1"
)

func (s *CartApi) AddItem(ctx context.Context, req *cart_v1.AddItemRequest) (*cart_v1.AddItemResponse, error) {

	cart, err := s.CartService.AddItem(req.GetCartId(), model.Item{
		Product_ID: req.GetProductId(),
		Quantity:   req.GetQuantity(),
	})
	if err != nil {
		return nil, err
	}

	fmt.Println(cart_v1.AddItemResponse{Cart: converter.FromModelCart(cart)})
	return &cart_v1.AddItemResponse{Cart: converter.FromModelCart(cart)}, nil
}
