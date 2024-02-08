package requests

type RegisterRequest struct {
	Name string `json:"name" binding:"required,min=1"`
	Email string `json:"email" binding:"required,email,min=1"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email  		string `json:"email"  validate:"required,email"`
	Password    string `json:"password"  validate:"required"`
}