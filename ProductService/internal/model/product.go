package model

import "encoding/json"

type Product struct {
	ID        int64   `json:"id" gorm:"primarykey" redis:"id"`
	Name      string  `json:"name" redis:"name"`
	Category  string  `json:"category" redis:"category"`
	Price     float64 `json:"price" redis:"price"`
	Inventory int64   `json:"inventory" redis:"inventory"`
}

func (u Product) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func (u *Product) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &u)
}