package repository

import (
	"context"
	"database/sql"
	"project-workshop/go-api-ecom/model/domain"
)

type OrderRepository interface {
	FindById(ctx context.Context, tx *sql.Tx, orderId int) (domain.Order, error)
	FindByUserId(ctx context.Context, tx *sql.Tx, userId int) ([]domain.Order, error)
	Insert(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order
	Update(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order
}