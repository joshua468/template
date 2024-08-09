package models

import (
    "time"
)

type User struct {
    ID        uint      `json:"id" gorm:"primary_key"`
    Username  string    `json:"username"`
    Email     string    `json:"email"`
    Password  string    `json:"password"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
