package helper

import (
	"project-workshop/go-api-ecom/model/domain"
	"project-workshop/go-api-ecom/model/web"
	// "strconv"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:       category.Id,
		Category: category.Category,
		Icon:     category.Icon,
		Products: ToProductResponses(category.Products),
	}
}

func ToUserResponse(user domain.User) web.UserResponse {
	role := "user"
	if !user.Role {
		role = "admin"
	}

	return web.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Role:     role,
	}
}

func ToAccountResponse(user domain.User) web.AccountResponse {
	role := "user"
	if !user.Role {
		role = "admin"
	}

	return web.AccountResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Role:     role,
	}
}

func ToProductResponse(product domain.Product) web.ProductResponse {
	return web.ProductResponse{
		Id:          product.Id,
		Name:        product.Name,
		Image:       product.Image,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity,
		CategoryId:  product.CategoryId,
		Category:    product.Category.Category,
	}
}

func ToOrderResponse(order domain.Order) web.OrderResponse {
	return web.OrderResponse{
		ID:            order.ID,
		UserID:        order.UserID,
		OrderItem:     ToCartResponses(order.OrderItems),
		TotalItems:    order.TotalItems,
		TotalPrice:    order.TotalPrice,
		OrderStatus:   string(order.OrderStatus),
		PaymentStatus: string(order.PaymentStatus),
	}
}

func ToCartResponse(cart domain.Cart) web.CartResponse {
	return web.CartResponse{
		Id:       cart.Id,
		UserId:   cart.UserId,
		Product:  ToProductResponses(cart.Product),
		Quantity: cart.Quantity,
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

func ToAccountResponses(users []domain.User) []web.AccountResponse {
	var accountResponses []web.AccountResponse
	for _, user := range users {
		accountResponses = append(accountResponses, ToAccountResponse(user))
	}
	return accountResponses
}

func ToProductResponses(products []domain.Product) []web.ProductResponse {
	var productResponses []web.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}
	return productResponses
}

func ToOrderResponses(orders []domain.Order) []web.OrderResponse {
	var orderResponses []web.OrderResponse
	for _, order := range orders {
		orderResponses = append(orderResponses, ToOrderResponse(order))
	}
	return orderResponses
}

func ToCartResponses(carts []domain.Cart) []web.CartResponse {
	var cartResponses []web.CartResponse
	for _, cart := range carts {
		cartResponses = append(cartResponses, ToCartResponse(cart))
	}
	return cartResponses
}
