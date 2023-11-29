package helper

import (
	"project-workshop/go-api-ecom/model/domain"
	"project-workshop/go-api-ecom/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:       category.Id,
		Category: category.Category,
		Products: ToProductResponses(category.Products),
	}
}

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Role:     user.Role,
	}
}

func ToProductResponse(product domain.Product) web.ProductResponse {
	return web.ProductResponse{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity,
		CategoryId:  product.CategoryId,
		Category:    product.Category.Category,
	}
}

func ToCartResponse(cart domain.Cart) web.CartResponse {
	return web.CartResponse{
		Id:        cart.Id,
		UserId:    cart.UserId,
		ProductId: cart.ProductId,
		Quantity:  cart.Quantity,
		Product:   ToProductResponse(cart.Product),
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}

func ToUserResponses(users []domain.User) []web.UserResponse {
	var userResponses []web.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}

func ToProductResponses(products []domain.Product) []web.ProductResponse {
	var productResponses []web.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}
	return productResponses
}

func ToCartResponses(carts []domain.Cart) []web.CartResponse {
	var cartResponses []web.CartResponse
	for _, cart := range carts {
		cartResponses = append(cartResponses, ToCartResponse(cart))
	}
	return cartResponses
}
