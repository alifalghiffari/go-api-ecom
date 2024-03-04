package web

type OrderCreateRequest struct {
	CartId []int `json:"cart"`
}

type OrderUpdateRequest struct {
	ID            int    `json:"id"`
	UserID        int    `json:"user_id"`
	TotalItems    int    `json:"total_items"`
	TotalPrice    int    `json:"total_price"`
	OrderStatus   string `json:"order_status"`
	PaymentStatus string `json:"payment_status"`
}
