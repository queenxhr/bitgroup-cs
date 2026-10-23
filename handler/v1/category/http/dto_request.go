package category

import (
	"bitgroup-cs/internal/entity"
)

type CreateCategoryRequestDto struct {
	Name string `json:"name" binding:"required"`
}

func (dto *CreateCategoryRequestDto) ToEntity() *entity.Category {
	if dto == nil {
		return nil
	}
	return &entity.Category{
		Name: dto.Name,
	}
}

type UpdateCategoryRequestDto struct {
	Name string `json:"name" binding:"required"`
}

func (dto *UpdateCategoryRequestDto) ToEntity(id int64) *entity.Category {
	if dto == nil {
		return nil
	}
	return &entity.Category{
		ID:   id,
		Name: dto.Name,
	}
}
