package service

import (
	"crudSystem/internal/entities"
	"fmt"

	"gorm.io/gorm"
)

func NewProduct(p *entities.ProductPGDB, db *gorm.DB) error {
	result := db.Create(p)
	fmt.Println(result)
	return nil
}
