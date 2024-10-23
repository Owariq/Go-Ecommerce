package api

import (
	"context"

	"github.com/Owariq/Go-Ecommerce/product-service/pkg/product_v1"
)

func (a *ProductApi) Delete(ctx context.Context, req *product_v1.DeleteRequest) (*product_v1.DeleteResponse, error) {

	id, err := a.ProductService.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &product_v1.DeleteResponse{Id: id}, nil
}