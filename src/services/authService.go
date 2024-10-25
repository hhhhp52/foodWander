package services

import (
	"database/sql"
	"errors"
	"foodWander/src/database"
	"foodWander/src/models"
	"foodWander/src/modules/authModule"
	"foodWander/src/utils/helpers"
	"net/http"
)

func Login(input authModule.LoginInput) (statusCode int, message string, data map[string]interface{}) {
	input.Password = helpers.TransferEncryptPassword(input.Password)
	tx, _ := database.DB().Begin()
	var user models.User
	err := tx.QueryRow(
		"SELECT * FROM user WHERE email = $1 AND password = $2", input.Email, input.Password).Scan(
		&user)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return http.StatusBadRequest, "The user isn't exist", nil
		}
		return http.StatusBadRequest, "Login isn't success", nil
	}

	var userToken models.UserToken

	err = tx.QueryRow("SELECT * FROM user_token WHERE user_id = $1", user.ID).Scan(&userToken)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return http.StatusInternalServerError, "Failed to login", nil
	}

	return http.StatusOK, "Login successful", nil
}

func Logout(input authModule.LogoutInput) (statusCode int, message string, data map[string]interface{}) {
	tx, _ := database.DB().Begin()
	_, err := tx.Exec("DELETE FROM user_token WHERE user_id = (SELECT id FROM user WHERE email = $1)", input.Email)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return http.StatusInternalServerError, "Failed to logout", nil
		}
		return http.StatusInternalServerError, "Failed to logout", nil
	}
	return http.StatusOK, "Logout successful", nil
}

func Register(input authModule.RegisterInput) (statusCode int, message string, data map[string]interface{}) {
	// Check email
	if !helpers.IsValidEmail(input.Email) {
		return http.StatusBadRequest, "Invalid email format", nil
	}

	// Check password
	if !helpers.IsValidPassword(input.Password) {
		return http.StatusBadRequest, "Password must be 6-16 characters long, contain one uppercase, one lowercase, and one special character", nil
	}

	// Check password confirmation
	if input.Password != input.ConfirmPassword {
		return http.StatusBadRequest, "Passwords do not match", nil
	}

	verifiedCode := helpers.GenerateVerificationCode()

	// Set password to encrypted password
	input.Password = helpers.TransferEncryptPassword(input.Password)

	tx, _ := database.DB().Begin()

	_, err := tx.Exec(
		"INSERT INTO user (email, password, catrgory, level, verified, verified_code) VALUES ($1, $2, $3, $4, $5, $6)",
		input.Email, input.Password, "user", "entry", false, verifiedCode)

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

	// Send verification code to email
	if err = helpers.SendVerificationEmail(input.Email, verifiedCode); err != nil {
		return http.StatusInternalServerError, "Failed to send verification email", nil
	}

	return http.StatusCreated, "User registered successfully", nil
}

func VerifyEmail(input authModule.VerifyEmailInput) (statusCode int, message string, data map[string]interface{}) {
	tx, _ := database.DB().Begin()
	var user models.User
	err := tx.QueryRow("SELECT * FROM user WHERE email = $1 AND verified_code = $2", input.Email).Scan(&user)

	if err != nil {
		return http.StatusBadRequest, "Invalid email", nil

	}

	if user.VerifiedCode != input.VerificationCode {
		return http.StatusBadRequest, "Invalid verification code", nil
	}

	_, err = tx.Exec("UPDATE user SET verified = $1 WHERE email = $2", true, input.Email)

	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return http.StatusInternalServerError, "Failed to verify email", nil
		}
		return http.StatusInternalServerError, "Failed to verify email", nil
	}

	return http.StatusOK, "Email verified successfully", nil
}

func ForgetPassword(input authModule.ForgetPasswordInput) (statusCode int, message string, data map[string]interface{}) {
	verifiedCode := helpers.GenerateVerificationCode()
	// Send verification code to email
	if err := helpers.SendVerificationEmail(input.Email, verifiedCode); err != nil {
		return http.StatusInternalServerError, "Failed to send verification email", nil
	}

	tx, _ := database.DB().Begin()
	_, err := tx.Exec("UPDATE user SET verified_code = $1, verified = false WHERE email = $2", verifiedCode, input.Email)

	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return http.StatusInternalServerError, "Failed to forget password", nil
		}
		return http.StatusInternalServerError, "Failed to forget password", nil
	}

	return http.StatusOK, "Forget password successful", nil
}

func ResetPassword(input authModule.ResetPasswordInput) (statusCode int, message string, data map[string]interface{}) {
	tx, _ := database.DB().Begin()

	var user models.User

	// Check password
	if !helpers.IsValidPassword(input.Password) {
		return http.StatusBadRequest, "Password must be 6-16 characters long, contain one uppercase, one lowercase, and one special character", nil
	}

	err := tx.QueryRow("SELECT * FROM user WHERE email = $1 ", input.Email).Scan(&user)
	if err != nil {
		return http.StatusBadRequest, "Invalid email", nil
	}

	if user.VerifiedCode != input.VerificationCode {
		return http.StatusBadRequest, "Invalid verification code", nil
	}

	// Set password to encrypted password
	input.Password = helpers.TransferEncryptPassword(input.Password)

	_, err = tx.Exec("UPDATE user SET password = $1 WHERE email = $2", input.Password, input.Email)

	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return http.StatusInternalServerError, "Failed to reset password", nil
		}
		return http.StatusInternalServerError, "Failed to reset password", nil
	}

	return http.StatusOK, "Reset password successful", nil
}
