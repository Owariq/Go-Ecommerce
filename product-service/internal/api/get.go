package api

import (
	"context"

	"github.com/Owariq/Go-Ecommerce/product-service/internal/utils/converter"
	"github.com/Owariq/Go-Ecommerce/product-service/pkg/product_v1"
)

func (a *ProductApi) Get(ctx context.Context, req *product_v1.GetRequest) (*product_v1.GetResponse, error) {

	product, err := a.ProductService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &product_v1.GetResponse{Product: converter.ModelToProto(product)}, nil
}