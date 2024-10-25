package authModule

type RegisterInput struct {
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LogoutInput struct {
	Email string `json:"email" binding:"required"`
}

type VerifyEmailInput struct {
	Email            string `json:"email" binding:"required"`
	VerificationCode string `json:"verification_code" binding:"required"`
}

type ForgetPasswordInput struct {
	Email string `json:"email" binding:"required"`
}

type ResetPasswordInput struct {
	Email            string `json:"email" binding:"required"`
	Password         string `json:"password" binding:"required"`
	VerificationCode string `json:"verification_code" binding:"required"`
}
