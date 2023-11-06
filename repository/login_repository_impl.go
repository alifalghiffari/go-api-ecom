package repository

import (
	"database/sql"
	"context"
	"project-workshop/go-api-ecom/model/domain"
	"project-workshop/go-api-ecom/helper"
)

type LoginRepositoryImpl struct {
}

func NewLoginRepository() LoginRepository {
	return &LoginRepositoryImpl{}
}

func (repository *LoginRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, username string) domain.User {
	SQL := "select id, username, password, email, role from user where username = ?"
	rows, err := tx.QueryContext(ctx, SQL, username)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Role)
		helper.PanicIfError(err)
		return user
	} else {
		return user
	}
}