package api

import (
	"context"

	"github.com/Owariq/Go-Ecommerce/product-service/internal/utils/converter"
	"github.com/Owariq/Go-Ecommerce/product-service/pkg/product_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *ProductApi) List(ctx context.Context, e *emptypb.Empty) (*product_v1.ListResponse, error) {
	products, err := a.ProductService.List(ctx)
	if err != nil {
		return nil, err
	}
	return &product_v1.ListResponse{Products: converter.ModelToProtoList(products)}, nil
}