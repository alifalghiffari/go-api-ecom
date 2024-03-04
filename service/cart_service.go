package service

import (
	"context"
	"project-workshop/go-api-ecom/model/web"
)

type CartService interface {
	AddToCart(ctx context.Context, request web.CartCreateRequest, userId int) web.CartResponse
	UpdateCart(ctx context.Context, request web.CartUpdateRequest, userId int) web.CartResponse
	DeleteCart(ctx context.Context, request web.CartDeleteRequest, userId int) web.CartResponse
	FindById(ctx context.Context, cartId int) web.CartResponse
	FindByUserId(ctx context.Context, userId int) []web.CartResponse
	FindAll(ctx context.Context) []web.CartResponse
}