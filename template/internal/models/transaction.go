package models

import (
    "time"
)

type Transaction struct {
    ID         uint      `json:"id" gorm:"primary_key"`
    UserID     uint      `json:"user_id"`
    ProductID  uint      `json:"product_id"`
    Quantity   int       `json:"quantity"`
    TotalPrice float64   `json:"total_price"`
    CreatedAt  time.Time `json:"created_at"`
    UpdatedAt  time.Time `json:"updated_at"`
}
