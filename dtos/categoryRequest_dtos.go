package dtos

type CategoryRequestDtos struct {
	Id   uint
	Name string `validate:"required,max=200,min=1"`
}
