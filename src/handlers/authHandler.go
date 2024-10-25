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

	statusCode, message, data := services.Register(c, input)
	c.JSON(statusCode, gin.H{"message": message, "data": data})
	return
}

func Login(c *gin.Context) {
	var input authModule.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

	if input.Email == "" || input.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

}

func VerifyEmail(c *gin.Context) {
	var input authModule.VerifyEmailInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

}

func Logout(c *gin.Context) {

}

func ForgetPassword(c *gin.Context) {

}

func ResetPassword(c *gin.Context) {

}
