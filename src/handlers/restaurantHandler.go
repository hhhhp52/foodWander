package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func GetRestaurants(c *gin.Context) {

}

func CreateRestaurant(c *gin.Context) {

}

func UpdateRestaurant(c *gin.Context) {

}

func DeleteRestaurant(c *gin.Context) {

}

func GetRestaurantDetail(c *gin.Context) {

	var input struct {
		RestaurantID uuid.UUID `json:"restaurant_id"`
	}

	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	log.Println(input.RestaurantID)
	c.String(200, "Success")
}
