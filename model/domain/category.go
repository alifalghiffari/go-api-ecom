package domain

type Category struct {
	Id       int
	Category string
	Icon     string
	Products []Product
}
