package services

import (
	"example/ecomers/models"
)

func GetBrands() []models.Brand {
	var brands []models.Brand
	DB.Find(&brands)

	return brands
}
