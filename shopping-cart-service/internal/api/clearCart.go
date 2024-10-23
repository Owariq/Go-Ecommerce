package api

import (
	"context"

	"github.com/Owariq/Go-Ecommerce/shopping-cart-service/pkg/cart_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *CartApi) ClearCart(ctx context.Context, req *cart_v1.ClearCartRequest) (*emptypb.Empty, error) {

	err := s.CartService.ClearCart(req.GetCartId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
