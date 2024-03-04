package repository

import (
	"context"
	"database/sql"
	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/model/domain"
)

type CartRepositoryImpl struct {
}

func NewCartRepositoryImpl() CartRepository {
	return &CartRepositoryImpl{}
}

func (repository *CartRepositoryImpl) AddToCart(ctx context.Context, tx *sql.Tx, cart domain.Cart) domain.Cart {
	SQL := "insert into cart(userId, product_id, quantity) values (?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, cart.UserId, cart.ProductId, cart.Quantity)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	cart.Id = int(id)
	return cart
}

func (repository *CartRepositoryImpl) UpdateCart(ctx context.Context, tx *sql.Tx, cart domain.Cart) domain.Cart {
	SQL := "update cart set quantity = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, cart.Quantity, cart.Id)
	helper.PanicIfError(err)

	return cart
}

func (repository *CartRepositoryImpl) DeleteCart(ctx context.Context, tx *sql.Tx, cart domain.Cart) domain.Cart {
	SQL := "delete from cart where id = ?"
	_, err := tx.ExecContext(ctx, SQL, cart.Id)
	helper.PanicIfError(err)

	return cart
}

// FindById fetches cart items based on a slice of cart IDs
func (repository *CartRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, cartId []int) ([]domain.Cart, error) {
	SQL := `
		SELECT c.id, c.userId, c.product_id, c.quantity, p.name, p.price
		FROM cart c
		INNER JOIN product p ON c.product_id = p.id
		WHERE c.id IN (?)
	`
	rows, err := tx.QueryContext(ctx, SQL, cartId)
	helper.PanicIfError(err)

	var carts []domain.Cart
	for rows.Next() {
		cart := domain.Cart{}
		product := domain.Product{}
		err := rows.Scan(&cart.Id, &cart.UserId, &cart.ProductId, &cart.Quantity, &product.Name, &product.Price)
		helper.PanicIfError(err)
		cart.Product = append(cart.Product, product)
		carts = append(carts, cart)
	}

	return carts, nil
}

func (repository *CartRepositoryImpl) FindByUserId(ctx context.Context, tx *sql.Tx, userId int) ([]domain.Cart, error) {
	SQL := `
		SELECT c.id, c.userId, c.product_id, c.quantity, p.name, p.price
		FROM cart c
		LEFT JOIN product p ON c.product_id = p.id
		WHERE c.userId = ?
    `
	rows, err := tx.QueryContext(ctx, SQL, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var carts []domain.Cart
	for rows.Next() {
		cart := domain.Cart{}
		product := domain.Product{}
		if err := rows.Scan(&cart.Id, &cart.UserId, &product.Id, &cart.Quantity, &product.Name, &product.Price); err != nil {
			return nil, err
		}
		cart.Product = append(cart.Product, product)
		carts = append(carts, cart)
	}

	return carts, nil
}

func (repository *CartRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Cart {
	SQL := "select id, userId, product_id, quantity from cart"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var carts []domain.Cart
	for rows.Next() {
		cart := domain.Cart{}
		err := rows.Scan(&cart.Id, &cart.UserId, &cart.ProductId, &cart.Quantity)
		helper.PanicIfError(err)
		carts = append(carts, cart)
	}

	return carts
}
