package service

import (
	"context"
	"project-workshop/go-api-ecom/model/web"
)

type LoginService interface {
	Login(ctx context.Context, tx *string, login web.LoginRequest) web.LoginResponse
}