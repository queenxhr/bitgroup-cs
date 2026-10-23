package category

import (
	"time"

	"bitgroup-cs/internal/entity"
)

type CategoryModel struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement"`
	Name      string    `gorm:"column:name;type:varchar(100);not null"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
}

func (CategoryModel) TableName() string {
	return "categories"
}

func (model CategoryModel) ToEntity() *entity.Category {
	return &entity.Category{
		ID:        model.ID,
		Name:      model.Name,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func (CategoryModel) FromEntity(ent *entity.Category) *CategoryModel {
	if ent == nil {
		return nil
	}
	return &CategoryModel{
		ID:        ent.ID,
		Name:      ent.Name,
		CreatedAt: ent.CreatedAt,
		UpdatedAt: ent.UpdatedAt,
	}
}

type listCategoryModel []CategoryModel

func (models listCategoryModel) ToEntities() []*entity.Category {
	entities := make([]*entity.Category, len(models))
	for i, item := range models {
		entities[i] = item.ToEntity()
	}
	return entities
}
