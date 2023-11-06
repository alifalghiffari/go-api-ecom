package domain

type Login struct {
	Username string `validate:"required,min=1,max=100" json:"username"`
	Password string `validate:"required,min=1,max=100" json:"password"`
}