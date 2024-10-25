package helpers

import (
	"crypto/sha256"
	"fmt"
	"foodWander/src/utils"
	"math/rand"
	"net/smtp"
	"regexp"
	"time"
)

func IsValidEmail(email string) bool {
	// Simple regex for email validation
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}

func IsValidPassword(password string) bool {
	// Password must be 6-16 characters long, contain one uppercase, one lowercase, and one special character
	if len(password) < 6 || len(password) > 16 {
		return false
	}
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`).MatchString(password)
	return hasLower && hasUpper && hasSpecial
}

func TransferPassword(password string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
}

func GenerateVerificationCode() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func SendVerificationEmail(email, code string) error {
	from := utils.SendEmail
	password := utils.SendEmailPassword
	to := email
	smtpHost := utils.SmtpHost
	smtpPort := utils.SmtpPort

	message := fmt.Sprintf("Subject: Email Verification Code\n\nYour verification code is: %s", code)

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(message))
	if err != nil {
		return err
	}
	return nil
}
