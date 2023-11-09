package web

type UserCreateRequest struct {
	Username string `json:"username" validate:"required,max=200,min=1"`
	Password string `json:"password" validate:"required,max=200,min=1"`
	Email    string `json:"email" validate:"required,max=200,min=1"`
	Role     string `json:"role" validate:"required,max=200,min=1"`
}