package web

type OrderResponse struct {
	ID            int            `json:"id"`
	UserID        int            `json:"user_id"`
	OrderItem     []CartResponse `json:"order_item"`
	TotalItems    int            `json:"total_items"`
	TotalPrice    int            `json:"total_price"`
	OrderStatus   string         `json:"order_status"`
	PaymentStatus string         `json:"payment_status"`
}
