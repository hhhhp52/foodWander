package handlers

import (
	"foodWander/src/modules/authModule"
	"foodWander/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	var input authModule.RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload", "data": nil})
		return
	}

	statusCode, message, data := services.Register(input)

	c.JSON(statusCode, gin.H{"message": message, "data": data})
	return
}

func Login(c *gin.Context) {
	var input authModule.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

	statusCode, message, data := services.Login(input)

	c.JSON(statusCode, gin.H{"message": message, "data": data})
	return

}

func VerifyEmail(c *gin.Context) {
	var input authModule.VerifyEmailInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

	statusCode, message, data := services.VerifyEmail(input)

	c.JSON(statusCode, gin.H{"message": message, "data": data})
	return

}

func Logout(c *gin.Context) {
	var input authModule.LogoutInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

	statusCode, message, data := services.Logout(input)

	c.JSON(statusCode, gin.H{"message": message, "data": data})
	return
}

func ForgetPassword(c *gin.Context) {
	var input authModule.ForgetPasswordInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

	statusCode, message, data := services.ForgetPassword(input)

	c.JSON(statusCode, gin.H{"message": message, "data": data})
	return
}

func ResetPassword(c *gin.Context) {
	var input authModule.ResetPasswordInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

	statusCode, message, data := services.ResetPassword(input)

	c.JSON(statusCode, gin.H{"message": message, "data": data})
	return
}
