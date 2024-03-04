package domain

type Payment struct {
	Id          int
	UserId      int
	OrderId     int
	Order       Order
	TotalAmount int
}
