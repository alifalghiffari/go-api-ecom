package app

import (
	"database/sql"
	"project-workshop/go-api-ecom/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/pakthani_db")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
