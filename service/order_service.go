package service

import (
	"context"
	"project-workshop/go-api-ecom/model/web"
)

type OrderService interface {
	CreateOrder(ctx context.Context, request web.OrderCreateRequest, userId int) web.OrderResponse
	UpdateOrder(ctx context.Context, request web.OrderUpdateRequest, Id int, userId int) web.OrderResponse
	FindOrderByUserId(ctx context.Context, userId int) []web.OrderResponse
	FindOrderById(ctx context.Context, Id int, userId int) web.OrderResponse
}

