package models

import "github.com/google/uuid"

type RestaurantRating struct {
	ID             int       `json:"id"`
	RestaurantID   uuid.UUID `json:"restaurant_id"`
	UserID         uuid.UUID `json:"user_id"`
	OverviewRating float32   `json:"overview_rating"`
	CpRating       float32   `json:"cp_rating"`
	MealRating     float32   `json:"meal_rating"`
	DrinkRating    float32   `json:"drink_rating"`
	ServiceRating  float32   `json:"service_rating"`
	AmbienceRating float32   `json:"ambience_rating"`
	Comment        string    `json:"comment"`
	CreatedAt      string    `json:"created_at"`
	UpdatedAt      string    `json:"updated_at"`
	DeletedAt      string    `json:"deleted_at"`
}

type RestaurantRatingHistory struct {
	ID                  int       `json:"id"`
	RestaurantID        uuid.UUID `json:"restaurant_id"`
	Month               string    `json:"month"`
	Year                string    `json:"year"`
	OverviewRatingAvg   float32   `json:"overview_rating_avg"`
	OverviewRatingCount float32   `json:"overview_rating_count"`
	CpRatingAvg         float32   `json:"cp_rating_avg"`
	CpRatingCount       float32   `json:"cp_rating_count"`
	MealRatingAvg       float32   `json:"meal_rating_avg"`
	MealRatingCount     float32   `json:"meal_rating_count"`
	DrinkRatingAvg      float32   `json:"drink_rating_avg"`
	DrinkRatingCount    float32   `json:"drink_rating_count"`
	ServiceRatingAvg    float32   `json:"service_rating_avg"`
	ServiceRatingCount  float32   `json:"service_rating_count"`
	AmbienceRatingAvg   float32   `json:"ambience_rating_avg"`
	AmbienceRatingCount float32   `json:"ambience_rating_count"`
	CreatedAt           string    `json:"created_at"`
	UpdatedAt           string    `json:"updated_at"`
	DeletedAt           string    `json:"deleted_at"`
}
