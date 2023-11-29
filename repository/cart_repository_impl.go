package repository

import (
	"context"
	"database/sql"
	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/model/domain"
)

type CartRepositoryImpl struct {
}

func NewCartRepository() CartRepository {
	return &CartRepositoryImpl{}
}

func (repository *CartRepositoryImpl) AddToCart(ctx context.Context, tx *sql.Tx, cart domain.Cart) domain.Cart {
	SQL := "insert into cart(user_id, product_id, quantity) values (?, ?, ?)"
	_, err := tx.ExecContext(ctx, SQL, cart.UserId, cart.ProductId, cart.Quantity)
	helper.PanicIfError(err)

	return cart
}

func (repository *CartRepositoryImpl) UpdateCart(ctx context.Context, tx *sql.Tx, cart domain.Cart) domain.Cart {
	SQL := "update cart set quantity = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, cart.Quantity, cart.Id)
	helper.PanicIfError(err)

	return cart
}

func (repository *CartRepositoryImpl) RemoveFromCart(ctx context.Context, tx *sql.Tx, cart domain.Cart) {
	SQL := "delete from cart where user_id = ? and product_id = ?"
	_, err := tx.ExecContext(ctx, SQL, cart.UserId, cart.ProductId)
	helper.PanicIfError(err)
}

func (repository *CartRepositoryImpl) GetItemsInCart(ctx context.Context, tx *sql.Tx, userId int) []domain.Cart {
	SQL := `
		SELECT c.id, c.user_id, c.product_id, c.quantity, p.name AS product_name, p.price AS product_price
		FROM cart c
		LEFT JOIN product p ON c.product_id = p.id
		WHERE c.user_id = ?
	`

	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)

	var carts []domain.Cart
	for rows.Next() {
		cart := domain.Cart{}
		err := rows.Scan(&cart.Id, &cart.UserId, &cart.ProductId, &cart.Quantity, &cart.Product.Name, &cart.Product.Price)
		helper.PanicIfError(err)
		carts = append(carts, cart)
	}

	return carts
}

func (repository *CartRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) domain.Cart {
	SQL := `
		SELECT c.id, c.user_id, c.product_id, c.quantity, p.name AS product_name, p.price AS product_price
		FROM cart c
		LEFT JOIN product p ON c.product_id = p.id
		WHERE c.id = ?
	`

	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)

	cart := domain.Cart{}
	if rows.Next() {
		err := rows.Scan(&cart.Id, &cart.UserId, &cart.ProductId, &cart.Quantity, &cart.Product.Name, &cart.Product.Price)
		helper.PanicIfError(err)
	}

	return cart
}