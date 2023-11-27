package domain

type Product struct {
	Id          int
	Name        string
	Description string
	Price       int
	Quantity    int
	// CategoryId is a foreign key to
	CategoryId  int
	Category    Category
}