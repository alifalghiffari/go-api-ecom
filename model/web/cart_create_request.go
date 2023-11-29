package web

type CartCreateRequest struct {
	UserId    int `validate:"required" json:"user_id"`
	ProductId int `validate:"required" json:"product_id"`
	Quantity  int `validate:"required" json:"quantity"`
}