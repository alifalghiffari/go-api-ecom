package repository

import (
	"database/sql"
	"context"
	"project-workshop/go-api-ecom/model/domain"
)

type LoginRepository interface {
	FindByUsername(ctx context.Context, tx *sql.Tx, username string) domain.User
}