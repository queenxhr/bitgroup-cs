package product

import (
	"time"

	"bitgroup-cs/infrastructure/repository/category"
	"bitgroup-cs/internal/entity"
)

type ProductModel struct {
	ID         int64                   `gorm:"column:id;primaryKey;autoIncrement"`
	Name       string                  `gorm:"column:name;type:varchar(150);not null"`
	Price      float64                 `gorm:"column:price;type:numeric(12,2);not null;default:0.00"`
	CategoryID int64                   `gorm:"column:category_id;not null"`
	Category   *category.CategoryModel `gorm:"foreignKey:CategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt  time.Time               `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time               `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
}

func (ProductModel) TableName() string {
	return "products"
}

func (model ProductModel) ToEntity() *entity.Product {
	ent := &entity.Product{
		ID:         model.ID,
		Name:       model.Name,
		Price:      model.Price,
		CategoryID: model.CategoryID,
		CreatedAt:  model.CreatedAt,
		UpdatedAt:  model.UpdatedAt,
	}
	if model.Category != nil {
		ent.Category = model.Category.ToEntity()
	}
	return ent
}

func (ProductModel) FromEntity(ent *entity.Product) *ProductModel {
	if ent == nil {
		return nil
	}
	return &ProductModel{
		ID:         ent.ID,
		Name:       ent.Name,
		Price:      ent.Price,
		CategoryID: ent.CategoryID,
		CreatedAt:  ent.CreatedAt,
		UpdatedAt:  ent.UpdatedAt,
	}
}

type listProductModel []ProductModel

func (models listProductModel) ToEntities() []*entity.Product {
	entities := make([]*entity.Product, len(models))
	for i, item := range models {
		entities[i] = item.ToEntity()
	}
	return entities
}
