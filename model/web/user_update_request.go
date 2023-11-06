package web

type UserUpdateRequest struct {
	Id       int    `json:"id" validate:"required"`
	Username string `json:"username" validate:"required,min=1,max=100"`
	Password string `json:"password" validate:"required,min=1,max=100"`
	Email    string `json:"email" validate:"required,email"`
}