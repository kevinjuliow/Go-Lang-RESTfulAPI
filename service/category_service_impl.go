package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-playground/validator"
	"golang-restful-api/dtos"
	"golang-restful-api/helper"
	"golang-restful-api/models/domain"
	"golang-restful-api/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validator          *validator.Validate
}

// Create a constructor of CategoryService
func NewCategoryServiceImpl(categoryRepository repository.CategoryRepository, DB *sql.DB, validator *validator.Validate) CategoryService {
	return &CategoryServiceImpl{CategoryRepository: categoryRepository, DB: DB, Validator: validator}
}

func (categoryService CategoryServiceImpl) GETById(ctx context.Context, id uint) dtos.CategoryResponseDtos {
	tx, err := categoryService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, errNotFound := categoryService.CategoryRepository.FindById(ctx, tx, id)
	helper.PanicIfError(errNotFound)

	return helper.ToCategoryResponseDtos(category)
}

func (categoryService CategoryServiceImpl) GETAll(ctx context.Context) []dtos.CategoryResponseDtos {
	tx, err := categoryService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categoryList := categoryService.CategoryRepository.Findall(ctx, tx)
	var categoryListResponse []dtos.CategoryResponseDtos
	for _, category := range categoryList {
		categoryListResponse = append(categoryListResponse, helper.ToCategoryResponseDtos(category))
	}
	return categoryListResponse
}

func (categoryService CategoryServiceImpl) PUT(ctx context.Context, requestDtos dtos.CategoryRequestDtos) dtos.CategoryResponseDtos {
	err := categoryService.Validator.Struct(requestDtos)
	helper.PanicIfError(err)
	tx, err := categoryService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	//find if the id exists
	category, err := categoryService.CategoryRepository.FindById(ctx, tx, requestDtos.Id)
	helper.PanicIfError(err)

	//update the Name from user
	category.Name = requestDtos.Name

	//Implements Update
	category = categoryService.CategoryRepository.Update(ctx, tx, category)

	//return in Response Type
	return helper.ToCategoryResponseDtos(category)
}

func (categoryService CategoryServiceImpl) POST(ctx context.Context, requestDtos dtos.CategoryRequestDtos) dtos.CategoryResponseDtos {
	err := categoryService.Validator.Struct(requestDtos)
	helper.PanicIfError(err)
	tx, err := categoryService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Create a new category from requestDtos.
	category := domain.Category{
		Name: requestDtos.Name,
	}

	//implements save from repo
	category = categoryService.CategoryRepository.Save(ctx, tx, category)

	//Convert into Responses Type
	return helper.ToCategoryResponseDtos(category)
}

func (categoryService CategoryServiceImpl) DELETE(ctx context.Context, id uint) {
	tx, err := categoryService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, errNotFound := categoryService.CategoryRepository.FindById(ctx, tx, id)
	helper.PanicIfError(errNotFound)

	if !categoryService.CategoryRepository.Delete(ctx, tx, id) {
		panic(errors.New("not Found"))
	}
}
