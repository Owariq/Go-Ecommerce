package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Cart struct {
	Cart_ID int64 `json:"cart_id" gorm:"primarykey"`
	User_ID int64 `json:"name"`
	Items   Items `json:"items" gorm:"type:jsonb"`
}

type Item struct {
	Product_ID int64 `json:"product_id"`
	Quantity   int32 `json:"quantity"`
}

type Items []Item

func (c Items) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *Items) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Invalid data type for Items")
	}
	return json.Unmarshal(bytes, &c)
}
