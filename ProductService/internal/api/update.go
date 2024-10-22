package api

import (
	"context"

	"github.com/Owariq/Go-Ecommerce/product-service/internal/utils/converter"
	"github.com/Owariq/Go-Ecommerce/product-service/pkg/product_v1"
)

func (a *ProductApi) Update(ctx context.Context, req *product_v1.UpdateRequest) (*product_v1.UpdateResponse, error) {
	id, err := a.ProductService.Update(ctx, converter.ProtoToModel(req.GetProduct()))
	if err != nil {
		return nil, err
	}
	return &product_v1.UpdateResponse{Id: id}, nil
}