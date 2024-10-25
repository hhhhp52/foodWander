package authModule

type RegisterInput struct {
	Email           string `json:"email"`
	PhoneNumber     string `json:"phone_number"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogoutInput struct {
	Email string `json:"email"`
}

type VerifyEmailInput struct {
	Email            string `json:"email"`
	VerificationCode string `json:"verification_code"`
}
