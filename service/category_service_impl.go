package service

import (
	"context"
	"golang-restful-api/dtos"
)

type CategoryServiceImpl struct{}

func (categoryService CategoryServiceImpl) GETById(ctx context.Context, id uint) dtos.CategoryResponseDtos {
	//TODO implement me
	panic("implement me")
}

func (categoryService CategoryServiceImpl) GETAll(ctx context.Context) []dtos.CategoryResponseDtos {
	//TODO implement me
	panic("implement me")
}

func (categoryService CategoryServiceImpl) PUT(ctx context.Context, requestDtos dtos.CategoryRequestDtos) dtos.CategoryResponseDtos {
	//TODO implement me
	panic("implement me")
}

func (categoryService CategoryServiceImpl) POST(ctx context.Context, requestDtos dtos.CategoryRequestDtos) dtos.CategoryResponseDtos {
	//TODO implement me
	panic("implement me")
}

func (categoryService CategoryServiceImpl) DELETE(ctx context.Context, id uint) {
	//TODO implement me
	panic("implement me")
}
