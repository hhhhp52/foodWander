package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Registering User"})
			return
		}
	}()
}

func Login(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Registering User"})
			return
		}
	}()

	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if input.Email == "" || input.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

}

func VerifyEmail(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Registering User"})
			return
		}
	}()

	var input struct {
		Email string `json:"email"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

}

func Logout(c *gin.Context) {

}

func ForgetPassword(c *gin.Context) {

}

func ResetPassword(c *gin.Context) {

}
