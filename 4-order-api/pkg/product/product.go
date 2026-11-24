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

type ProductHandlerCfg struct {
	Repo *Repo
}

type ProductHandler struct {
	Repo *Repo
}

type ProductRequest struct {
	Name   string         `json:"name" validate:"required"`
	Desc   string         `json:"desc" validate:"required"`
	Price  int            `json:"price" validate:"required,gt=0"`
	Images pq.StringArray `json:"images" gorm:"type:text[]"`
}

type ProductUpdate struct {
	Name   string         `json:"name"`
	Desc   string         `json:"desc"`
	Price  int            `json:"price"`
	Images pq.StringArray `json:"images" gorm:"type:text[]"`
}
