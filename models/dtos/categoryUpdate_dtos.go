package dtos

type CategoryUpdateDtos struct {
	Id   uint   `validate:"required"`
	Name string `validate:"required max=200,min=1" json:"name"`
}
