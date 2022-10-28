package model

type AuthDTO struct {
	Uuid  string
	Email string
	Role  string
}

type SignInDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
