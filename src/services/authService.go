package services

import (
	"crypto/sha256"
	"fmt"
	"foodWander/src/database"
	"foodWander/src/modules/authModule"
	"foodWander/src/utils/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context, input authModule.LoginInput) (statusCode int, message string, data map[string]interface{}) {
	if input.Email == "" || input.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return http.StatusBadRequest, "Invalid request payload", nil
	}
	return http.StatusOK, "Login successful", nil
}

func Logout(c *gin.Context, input authModule.LogoutInput) (statusCode int, message string, data map[string]interface{}) {
	return http.StatusOK, "Logout successful", nil
}

func Register(c *gin.Context, input authModule.RegisterInput) (statusCode int, message string, data map[string]interface{}) {
	// Check email
	if !helpers.IsValidEmail(input.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	// Check password
	if !helpers.IsValidPassword(input.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be 6-16 characters long, contain one uppercase, one lowercase, and one special character"})
		return
	}

	// Check password confirmation
	if input.Password != input.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		return
	}

	verifiedCode := helpers.GenerateVerificationCode()

	// Send verification code to email
	if err := helpers.SendVerificationEmail(input.Email, verifiedCode); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send verification email"})
		return
	}

	// Set password to crypted password
	input.Password = fmt.Sprintf("%x", sha256.Sum256([]byte(input.Password)))

	tx, _ := database.DB().Begin()

	_, err := tx.Exec("INSERT INTO user (email, phone_number, password, verified_code) VALUES ($1, $2, $3, $4)",
		input.Email, input.PhoneNumber, input.Password, verifiedCode)

	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return http.StatusInternalServerError, "Failed to register user", nil
		}
		return http.StatusInternalServerError, "Failed to register user", nil
	}

	err = tx.Commit()
	if err != nil {
		return http.StatusInternalServerError, "Failed to register user", nil
	}

	return http.StatusCreated, "User registered successfully", nil
}

func EmailVerify(c *gin.Context, input authModule.VerifyEmailInput) {

}

func ForgetPassword(c *gin.Context) {

}

func ResetPassword(c *gin.Context) {

}
