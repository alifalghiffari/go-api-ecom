package repository

import (
	"context"
	"database/sql"
	"errors"
	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/model/domain"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "insert into product(name, price, category_id) values (?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, product.Name, product.Price, product.Category.Id)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	product.Id = int(id)
	return product
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "update product set name = ?, price = ?, category_id = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Name, product.Price, product.Category.Id, product.Id)
	helper.PanicIfError(err)

	return product
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product domain.Product) {
	SQL := "delete from product where id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Id)
	helper.PanicIfError(err)
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Product, error) {
	SQL := `
		SELECT p.id, p.name, p.image, p.description, p.price, p.category_id, c.category
		FROM product p
		JOIN category c ON p.category_id = c.id
		WHERE p.id = ?
	`
	rows, err := tx.QueryContext(ctx, SQL, productId)
	if err != nil {
		return domain.Product{}, err // Mengembalikan product kosong dan error
	}
	defer rows.Close()

	product := domain.Product{}
	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Name, &product.Image, &product.Description, &product.Price, &product.CategoryId, &product.Category.Category)
		if err != nil {
			return domain.Product{}, err // Mengembalikan product kosong dan error jika terjadi kesalahan saat pemindaian rows
		}
		return product, nil // Mengembalikan product yang ditemukan tanpa error
	}

	return domain.Product{}, errors.New("product is not found") // Jika tidak ada baris yang ditemukan, kembalikan error bahwa produk tidak ditemukan
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Product {
	SQL := `
		SELECT p.id, p.name, p.image, p.description, p.price, p.category_id, c.category
		FROM product p
		JOIN category c ON p.category_id = c.id
	`
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil // Mengembalikan nil slice product dan error
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		product := domain.Product{}
		err := rows.Scan(&product.Id, &product.Name, &product.Image, &product.Description, &product.Price, &product.CategoryId, &product.Category.Category)
		if err != nil {
			return nil // Mengembalikan nil slice product dan error jika terjadi kesalahan saat pemindaian rows
		}
		products = append(products, product) // Menambahkan product ke slice products
	}

	if len(products) == 0 {
		return nil // Jika tidak ada produk yang ditemukan, kembalikan error
	}

	return products // Mengembalikan slice product dan tanpa error
}
