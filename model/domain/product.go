package domain

type Product struct {
	Id          int
	Name        string
	Image       string
	Description string
	Price       int
	Quantity    int
	CategoryId  int
	Category    Category
}