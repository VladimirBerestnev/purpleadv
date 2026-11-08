package product

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name   string         `json:"name" validate:"required"`
	Desc   string         `json:"desc" validate:"required"`
	Price  int            `json:"price" validate:"required"`
	Images pq.StringArray `json:"images" gorm:"type:text[]"`
}
