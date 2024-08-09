package models

import (
    "time"
)

type Sales struct {
    ID          uint      `json:"id" gorm:"primary_key"`
    UserID      uint      `json:"user_id"`
    ProductID   uint      `json:"product_id"`
    Quantity    int       `json:"quantity"`
    TotalAmount float64   `json:"total_amount"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
