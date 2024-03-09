package helper

import (
	"golang-restful-api/models/domain"
	"golang-restful-api/models/dtos"
)

func ToCategoryResponseDtos(category domain.Category) dtos.CategoryResponseDtos {
	return dtos.CategoryResponseDtos{
		Id : category.Id,
		Name: category.Name,
	}
}
