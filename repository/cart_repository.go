package repository

import (
	"context"
	"database/sql"
	"project-workshop/go-api-ecom/model/domain"
)

type CartRepository interface {
	AddToCart(ctx context.Context, tx *sql.Tx, cart domain.Cart) domain.Cart
	UpdateCart(ctx context.Context, tx *sql.Tx, cart domain.Cart) domain.Cart
	RemoveFromCart(ctx context.Context, tx *sql.Tx, cart domain.Cart)
	GetItemsInCart(ctx context.Context, tx *sql.Tx, userId int) []domain.Cart
	FindById(ctx context.Context, tx *sql.Tx, id int) domain.Cart
}