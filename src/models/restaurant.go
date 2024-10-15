package models

import "github.com/google/uuid"

type Restaurant struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Address      string    `json:"address"`
	PhoneNumber  string    `json:"phone_number"`
	Email        string    `json:"email"`
	Website      string    `json:"website"`
	BelongToUser string    `json:"belong_to_user"`
	CreatedAt    string    `json:"created_at"`
	UpdatedAt    string    `json:"updated_at"`
	DeletedAt    string    `json:"deleted_at"`
}

type RestaurantMenu struct {
	ID             int       `json:"id"`
	RestaurantID   uuid.UUID `json:"restaurant_id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	MenuCategoryID int       `json:"menu_category_id"`
	Price          string    `json:"price"`
	Currency       string    `json:"currency"`
	CreatedAt      string    `json:"created_at"`
	UpdatedAt      string    `json:"updated_at"`
	DeletedAt      string    `json:"deleted_at"`
}

type MenuCategory struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
