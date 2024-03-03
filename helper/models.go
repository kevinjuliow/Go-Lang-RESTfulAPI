package helper

import (
	"golang-restful-api/dtos"
	"golang-restful-api/models/domain"
)

func ToCategoryResponseDtos(category domain.Category) dtos.CategoryResponseDtos {
	return dtos.CategoryResponseDtos{
		Name: category.Name,
	}
}
