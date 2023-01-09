package dto

type SecureCredentials struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
