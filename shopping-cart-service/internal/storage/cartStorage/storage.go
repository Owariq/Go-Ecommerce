package cartStorage

import (
	"fmt"

	"github.com/Owariq/Go-Ecommerce/shopping-cart-service/internal/model"
	"github.com/Owariq/Go-Ecommerce/shopping-cart-service/internal/storage"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewCartStorage(db *gorm.DB) storage.CartStorage {
	return &repo{
		db: db,
	}
}

func (r *repo) CreateCart(userID int64) (int64, error) {
	cart := model.Cart{User_ID: userID}
	err := r.db.Create(&cart).Error
	if err != nil {
		return 0, fmt.Errorf("error creating cart: %w", err)
	}
	return cart.Cart_ID, nil
}
func (r *repo) GetCart(cartID int64) (model.Cart, error) {
	var cart model.Cart
	err := r.db.First(&cart, cartID).Error
	if err != nil {
		return model.Cart{}, fmt.Errorf("error getting cart: %w", err)
	}
	return cart, nil
}

func (r *repo) AddItem(cartID int64, item model.Item) (model.Cart, error) {

	cart, err := r.GetCart(cartID)
	if err != nil {
		return model.Cart{}, fmt.Errorf("error getting cart: %w", err)
	}
	cart.Items = append(cart.Items, item)
	err = r.db.Save(&cart).Error
	if err != nil {
		return model.Cart{}, fmt.Errorf("error saving cart: %w", err)
	}
	return cart, nil
}
func (r *repo) RemoveItem(cartID int64, productID int64) (model.Cart, error) {

	cart, err := r.GetCart(cartID)
	if err != nil {
		return model.Cart{}, fmt.Errorf("error getting cart: %w", err)
	}
	for i, item := range cart.Items {
		if item.Product_ID == productID {
			cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
			err = r.db.Save(&cart).Error
			if err != nil {
				return model.Cart{}, fmt.Errorf("error saving cart: %w", err)
			}
			break
		}
	}
	return cart, nil
}
func (r *repo) ClearCart(cartID int64) error {

	cart, err := r.GetCart(cartID)
	if err != nil {
		return fmt.Errorf("error getting cart: %w", err)
	}
	cart.Items = []model.Item{}
	err = r.db.Save(&cart).Error
	if err != nil {
		return fmt.Errorf("error saving cart: %w", err)
	}
	return nil
}
