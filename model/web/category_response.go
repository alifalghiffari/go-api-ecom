package web

type CategoryResponse struct {
	Id       int               `json:"id"`
	Category string            `json:"category"`
	Icon     string            `json:"icon"`
	Products []ProductResponse `json:"products"`
}
