package api

import (
	"context"
	"log"

	"github.com/Owariq/Go-Ecommerce/product-service/internal/utils/converter"
	"github.com/Owariq/Go-Ecommerce/product-service/pkg/product_v1"
)

func (a *ProductApi) Create(ctx context.Context, req *product_v1.CreateRequest) (*product_v1.CreateResponse, error) {

	id, err := a.ProductService.Create(ctx, converter.ProtoToModel(req.GetProduct()))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted id: %d", id)

	
	return &product_v1.CreateResponse{
		Id: id,
	}, nil
}