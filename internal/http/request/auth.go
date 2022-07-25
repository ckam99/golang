package request

type RegisterRequest struct {
	Name     string `validate:"required,min=2,max=60"`
	Phone    string `validate:"min=6,max=60"`
	Email    string `validate:"required,email,max=255"`
	Password string `validate:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required"`
}
