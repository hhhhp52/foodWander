package models

import "github.com/google/uuid"

type RestaurantReview struct {
	ID             int       `json:"id"`
	RestaurantID   uuid.UUID `json:"restaurant_id"`
	UserID         uuid.UUID `json:"user_id"`
	FoodRating     float32   `json:"food_rating"`
	ServiceRating  float32   `json:"service_rating"`
	AmbienceRating float32   `json:"ambience_rating"`
	Comment        string    `json:"comment"`
}
