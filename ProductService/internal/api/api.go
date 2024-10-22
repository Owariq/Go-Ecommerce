package api

import (
	"github.com/Owariq/Go-Ecommerce/product-service/internal/service"
	"github.com/Owariq/Go-Ecommerce/product-service/pkg/product_v1"
)

type ProductApi struct {
	product_v1.UnimplementedProductV1Server
	ProductService service.ProductService
}

func NewProductApi(productService service.ProductService) *ProductApi {
	return &ProductApi{ProductService: productService}
}