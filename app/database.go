package app

import (
	"database/sql"
	"golang-restful-api/helper"
	"time"
)

func NewDb() *sql.DB {
	db, err := sql.Open("mysql", "root:ewk241xv28@tcp(localhost:3360)/golang_restful_api")
	helper.PanicIfError(err)

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
