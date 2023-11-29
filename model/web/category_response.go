package web

type CategoryResponse struct {
	Id       int    `json:"id"`
	Category string `json:"category"`
	Products []ProductResponse `json:"products"`
}
