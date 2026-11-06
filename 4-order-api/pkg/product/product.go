package product

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Price int    `json:"price"`
	Hash  string `json:"hash" gorm:"uniqueIndex"`
}
