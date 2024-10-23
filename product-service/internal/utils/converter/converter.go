package converter

import (
	"github.com/Owariq/Go-Ecommerce/product-service/internal/model"
	"github.com/Owariq/Go-Ecommerce/product-service/pkg/product_v1"
)

func ProtoToModel(product *product_v1.Product) model.Product {
	return model.Product{
		ID:        product.Id,
		Name:      product.Name,
		Category:  product.Category,
		Price:     product.Price,
		Inventory: product.Inventory,
	}
}

func ModelToProto(product model.Product) *product_v1.Product {
	return &product_v1.Product{
		Id:        product.ID,
		Name:      product.Name,
		Category:  product.Category,
		Price:     product.Price,
		Inventory: product.Inventory,
	}
}

func ModelToProtoList(products []model.Product) []*product_v1.Product {
	var list []*product_v1.Product
	for _, product := range products {
		list = append(list, ModelToProto(product))
	}
	return list
}