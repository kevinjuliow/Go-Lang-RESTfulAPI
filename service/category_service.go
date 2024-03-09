package service

import (
	"context"
	dtos2 "golang-restful-api/models/dtos"
)

type CategoryService interface {
	GETById(ctx context.Context, id uint) dtos2.CategoryResponseDtos
	GETAll(ctx context.Context) []dtos2.CategoryResponseDtos
	PUT(ctx context.Context, requestDtos dtos2.CategoryUpdateDtos) dtos2.CategoryResponseDtos
	POST(ctx context.Context, requestDtos dtos2.CategoryCreateDtos) dtos2.CategoryResponseDtos
	DELETE(ctx context.Context, id uint)
}
