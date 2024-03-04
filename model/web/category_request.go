package web

type CategoryCreateRequest struct {
	Category string `validate:"required,min=1,max=100" json:"name"`
}

type CategoryUpdateRequest struct {
	Id       int    `validate:"required"`
	Category string `validate:"required,max=200,min=1" json:"name"`
}

type CategoryDeleteRequest struct {
	Id int `validate:"required"`
}
