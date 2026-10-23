package category

import (
	"bitgroup-cs/internal/repository"
	"gorm.io/gorm"
)

// NewMysqlRepository aliases the repository constructor following logistic-service convention
func NewMysqlRepository(dbConn *gorm.DB) repository.CategoryRepository {
	return NewPostgresRepository(dbConn)
}
