package user

import (
	"errors"
	"order-api/pkg/db"

	"gorm.io/gorm/clause"
)

type UserRepository struct {
	database *db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{
		database: database,
	}
}

func (userRepo *UserRepository) Create(user *User) (*User, error) {
	result := userRepo.database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (userRepo *UserRepository) GetByPhone(phone string) (*User, error) {
	var user User
	result := userRepo.database.DB.First(&user, "phone = ?", phone)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (userRepo *UserRepository) GetBySessionID(sessionId string) (*User, error) {
	var user User
	result := userRepo.database.DB.First(&user, "session_id = ?", sessionId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (userRepo *UserRepository) Update(user *User) (*User, error) {
	result := userRepo.database.DB.Clauses(clause.Returning{}).Updates(user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}
	return user, nil
}
