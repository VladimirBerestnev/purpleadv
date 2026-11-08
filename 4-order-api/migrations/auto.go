package main

import (
	"order-api/pkg/product"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(".env don`t load")
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic("Error open connection bd")
	}
	errMigrate := db.AutoMigrate(&product.Product{})
	if errMigrate != nil {
		panic("Migration error")
	}
}
