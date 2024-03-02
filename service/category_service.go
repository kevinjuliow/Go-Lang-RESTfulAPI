package service

import (
	"context"
	"golang-restful-api/dtos"
)

type CategoryService interface {
	GETById(ctx context.Context, id uint) dtos.CategoryResponseDtos
	GETAll(ctx context.Context) []dtos.CategoryResponseDtos
	PUT(ctx context.Context, requestDtos dtos.CategoryRequestDtos) dtos.CategoryResponseDtos
	POST(ctx context.Context, requestDtos dtos.CategoryRequestDtos) dtos.CategoryResponseDtos
	DELETE(ctx context.Context, id uint)
}
