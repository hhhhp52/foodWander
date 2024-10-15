package routes

import (
	"foodWander/src/handlers"
)

func Router() {
	router := gin.Default()

	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)
	router.POST("/verify-email", handlers.VerifyEmail)
	router.POST("/logout", handlers.Logout)

}
