package web

type CartRemoveRequest struct {
	Id int `validate:"required" json:"id"`
	UserId int `validate:"required" json:"user_id"`
	ProductId int `validate:"required" json:"product_id"`
}