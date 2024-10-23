package converter

import (
	"github.com/Owariq/Go-Ecommerce/shopping-cart-service/internal/model"
	"github.com/Owariq/Go-Ecommerce/shopping-cart-service/pkg/cart_v1"
)

func FromModelCart(cart model.Cart) *cart_v1.Cart {
	return &cart_v1.Cart{
		CartId: cart.Cart_ID,
		UserId: cart.User_ID,
		Items:  FromModelItems(cart.Items),
	}
}

func FromModelItems(items []model.Item) []*cart_v1.Item {
	var res []*cart_v1.Item
	for _, item := range items {
		res = append(res, &cart_v1.Item{
			ProductId: item.Product_ID,
			Quantity:  item.Quantity,
		})
	}
	return res
}
