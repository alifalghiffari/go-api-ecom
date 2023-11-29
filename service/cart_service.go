package service

import (
	"context"
	"project-workshop/go-api-ecom/model/web"
)

type CartService interface {
	AddToCart(ctx context.Context, request web.CartCreateRequest) web.CartResponse
	UpdateCart(ctx context.Context, request web.CartUpdateRequest) web.CartResponse
	RemoveFromCart(ctx context.Context, request web.CartRemoveRequest)
	GetItemsInCart(ctx context.Context, userId int) []web.CartResponse
	FindById(ctx context.Context, id int) web.CartResponse
}