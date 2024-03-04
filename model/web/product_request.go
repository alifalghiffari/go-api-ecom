package web

type ProductCreateRequest struct {
	Name        string `json:"name" validate:"required,max=200,min=1"`
	Description string `json:"description" validate:"required,max=200,min=1"`
	Price       int    `json:"price" validate:"required,min=1"`
	Quantity    int    `json:"quantity" validate:"required,min=1"`
	CategoryId  int    `json:"categoryId" validate:"required"`
}

type ProductUpdateRequest struct {
	Id          int    `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required,max=200,min=1"`
	Description string `json:"description" validate:"required,max=200,min=1"`
	Price       int    `json:"price" validate:"required,min=1"`
	Quantity    int    `json:"quantity" validate:"required,min=1"`
	CategoryId  int    `json:"categoryId" validate:"required"`
}

type ProductDeleteRequest struct {
	Id int `json:"id" validate:"required"`
}
