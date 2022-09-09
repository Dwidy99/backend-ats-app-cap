package dto

// LoginDTO is model that used by client when POST from /login endpoint route
type LoginDTO struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}
