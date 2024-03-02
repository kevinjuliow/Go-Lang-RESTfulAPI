package repository

import (
	"context"
	"database/sql"
	"golang-restful-api/models/domain"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Findall(ctx context.Context, tx *sql.Tx) []domain.Category
	FindById(ctx context.Context, tx *sql.Tx, id uint) (domain.Category, error)
	Delete(ctx context.Context, tx *sql.Tx, id uint) bool
}
