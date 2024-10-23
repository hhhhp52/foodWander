package models

import "github.com/google/uuid"

type Fundraising struct {
	ID                 uuid.UUID `json:"id"`
	UserID             uuid.UUID `json:"user_id"`
	RestaurantID       uuid.UUID `json:"restaurant_id"`
	Title              string    `json:"title"`
	Status             string    `json:"status"`
	LowestTargetPrice  float32   `json:"lowest_target_price"`
	SuccessTargetPrice float32   `json:"success_target_price"`
	Currency           string    `json:"currency"`
	Introduction       string    `json:"introduction"`
	EndDate            string    `json:"end_date"`
	CreatedAt          string    `json:"created_at"`
	UpdatedAt          string    `json:"updated_at"`
	DeletedAt          string    `json:"deleted_at"`
}

type FundraisingSponsor struct {
	ID            int       `json:"id"`
	FundraisingID uuid.UUID `json:"fundraising_id"`
	UserID        uuid.UUID `json:"user_id"`
	Status        string    `json:"status"`
	Amount        float32   `json:"amount"`
	Currency      string    `json:"currency"`
	CreatedAt     string    `json:"created_at"`
	UpdatedAt     string    `json:"updated_at"`
	DeletedAt     string    `json:"deleted_at"`
}
