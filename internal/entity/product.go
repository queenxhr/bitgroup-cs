package entity

import "time"

type Product struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Price      float64   `json:"price"`
	CategoryID int64     `json:"category_id"`
	Category   *Category `json:"category,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
