package service

import (
	"context"
	"project-workshop/go-api-ecom/model/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse
	Delete(ctx context.Context, userId int)
	FindById(ctx context.Context, userId int) web.UserResponse
	FetchUserRole(ctx context.Context, role string) web.UserResponse
}