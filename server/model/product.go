package model

import (
	"server/database"
	"server/entity"
)

func FindAllProducts() []entity.Product {
	products := []entity.Product{}
	db := database.Connect()
	db.Order("ID asc").Find(&products)
	defer db.Close()

	return products
}
