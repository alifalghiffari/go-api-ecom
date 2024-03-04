package repository

import (
	"context"
	"database/sql"
	"project-workshop/go-api-ecom/model/domain"
)

type CartRepository interface {
	AddToCart(ctx context.Context, tx *sql.Tx, cart domain.Cart) domain.Cart
	UpdateCart(ctx context.Context, tx *sql.Tx, cart domain.Cart) domain.Cart
	DeleteCart(ctx context.Context, tx *sql.Tx, cart domain.Cart) domain.Cart
	FindById(ctx context.Context, tx *sql.Tx, cartId []int) ([]domain.Cart, error)
	FindByUserId(ctx context.Context, tx *sql.Tx, userId int) ([]domain.Cart, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Cart
}