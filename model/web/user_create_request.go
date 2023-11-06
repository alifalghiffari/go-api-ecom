package web

type UserCreateRequest struct {
	Username string `validate:"required,min=1,max=100" json:"username"`
	Password string `validate:"required,min=1,max=100" json:"password"`
	Email    string `validate:"required,email" json:"email"`
	Role	 string `validate:"required" json:"role"`
}