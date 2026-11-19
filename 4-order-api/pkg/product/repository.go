package product

import (
	"order-api/pkg/db"

	"gorm.io/gorm/clause"
)

type Repo struct {
	db *db.Db
}

func NewRepo(db *db.Db) *Repo {
	return &Repo{
		db: db,
	}
}

func (repo *Repo) Create(product *Product) (*Product, error) {
	result := repo.db.DB.Create(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (repo *Repo) GetById(id uint64) (*Product, error) {
	var product Product
	result := repo.db.DB.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (repo *Repo) Update(product *Product) (*Product, error) {
	result := repo.db.DB.Clauses(clause.Returning{}).Updates(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (repo *Repo) Delete(id uint64) error {
	result := repo.db.DB.Delete(&Product{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
