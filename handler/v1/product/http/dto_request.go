package product

import (
	"bitgroup-cs/internal/entity"
)

type CreateProductRequestDto struct {
	Name       string  `json:"name" binding:"required"`
	Price      float64 `json:"price" binding:"gte=0"`
	CategoryID int64   `json:"category_id" binding:"required,gt=0"`
}

func (dto *CreateProductRequestDto) ToEntity() *entity.Product {
	if dto == nil {
		return nil
	}
	return &entity.Product{
		Name:       dto.Name,
		Price:      dto.Price,
		CategoryID: dto.CategoryID,
	}
}

type UpdateProductRequestDto struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price" binding:"gte=0"`
	CategoryID int64   `json:"category_id"`
}

func (dto *UpdateProductRequestDto) ToEntity(id int64) *entity.Product {
	if dto == nil {
		return nil
	}
	return &entity.Product{
		ID:         id,
		Name:       dto.Name,
		Price:      dto.Price,
		CategoryID: dto.CategoryID,
	}
}
