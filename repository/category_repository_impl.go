package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-restful-api/helper"
	"golang-restful-api/models/domain"
)

type CategoryRepositoryImpl struct{}

// Create constructor of CategoryRepo
func NewCategoryRepositoryImpl() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repo *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	querry := "insert into category (name) values (?)"
	result, err := tx.ExecContext(ctx, querry, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	category.Id = uint(id)

	return category
}

func (repo *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	querry := "update category set name = ? where id = ? "
	_, err := tx.ExecContext(ctx, querry, category.Name, category.Id)
	helper.PanicIfError(err)
	return category
}

func (repo *CategoryRepositoryImpl) Findall(ctx context.Context, tx *sql.Tx) []domain.Category {
	querry := "select * from category"

	rows, err := tx.QueryContext(ctx, querry)
	helper.PanicIfError(err)

	var categoryList []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categoryList = append(categoryList, category)
	}

	return categoryList
}

func (repo *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id uint) (domain.Category, error) {
	querry := "select * from category where id = ?"
	rows, err := tx.QueryContext(ctx, querry, id)

	helper.PanicIfError(err)
	foundCategory := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&foundCategory.Id, &foundCategory.Name)
		helper.PanicIfError(err)
		return foundCategory, nil
	} else {
		return foundCategory, errors.New("Id Not Found")
	}
}

func (repo *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id uint) bool {
	querry := "delete from category where id = ?"
	_, err := tx.ExecContext(ctx, querry, id)
	helper.PanicIfError(err)
	return true
}
