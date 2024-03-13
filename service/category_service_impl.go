package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-playground/validator"
	"golang-restful-api/ErrorHandler"
	"golang-restful-api/helper"
	"golang-restful-api/models/domain"
	dtos2 "golang-restful-api/models/dtos"
	"golang-restful-api/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validator          *validator.Validate
}

// Create a constructor of CategoryService
func NewCategoryServiceImpl(categoryRepository repository.CategoryRepository, DB *sql.DB, validator *validator.Validate) *CategoryServiceImpl {
	return &CategoryServiceImpl{CategoryRepository: categoryRepository, DB: DB, Validator: validator}
}

func (categoryService *CategoryServiceImpl) GETById(ctx context.Context, id uint) dtos2.CategoryResponseDtos {
	tx, err := categoryService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, errNotFound := categoryService.CategoryRepository.FindById(ctx, tx, id)
	if errNotFound != nil {
		panic(ErrorHandler.NewNotFoundError(errNotFound.Error()))
	}

	return helper.ToCategoryResponseDtos(category)
}

func (categoryService *CategoryServiceImpl) GETAll(ctx context.Context) []dtos2.CategoryResponseDtos {
	tx, err := categoryService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categoryList := categoryService.CategoryRepository.Findall(ctx, tx)
	var categoryListResponse []dtos2.CategoryResponseDtos
	for _, category := range categoryList {
		categoryListResponse = append(categoryListResponse, helper.ToCategoryResponseDtos(category))
	}
	return categoryListResponse
}

func (categoryService *CategoryServiceImpl) PUT(ctx context.Context, requestDtos dtos2.CategoryUpdateDtos) dtos2.CategoryResponseDtos {
	err := categoryService.Validator.Struct(requestDtos)
	helper.PanicIfError(err)
	tx, err := categoryService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	//find if the id exists
	category, err := categoryService.CategoryRepository.FindById(ctx, tx, requestDtos.Id)
	if err != nil {
		panic(ErrorHandler.NotFoundError{err.Error()})
	}

	//update the Name from user
	category.Name = requestDtos.Name

	//Implements Update
	category = categoryService.CategoryRepository.Update(ctx, tx, category)

	//return in Response Type
	return helper.ToCategoryResponseDtos(category)
}

func (categoryService *CategoryServiceImpl) POST(ctx context.Context, requestDtos dtos2.CategoryCreateDtos) dtos2.CategoryResponseDtos {
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

func (categoryService *CategoryServiceImpl) DELETE(ctx context.Context, id uint) {
	tx, err := categoryService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, errNotFound := categoryService.CategoryRepository.FindById(ctx, tx, id)
	if errNotFound != nil {
		panic(ErrorHandler.NotFoundError{err.Error()})
	}

	if !categoryService.CategoryRepository.Delete(ctx, tx, id) {
		panic(errors.New("not Found"))
	}
}
