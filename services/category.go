package services

import "example/ecomers/models"

func GetCategories() []models.Category {

	var categories []models.Category

	DB.Model(&models.Category{}).Where("category_id IS NULL").Preload("Childrens").Find(&categories)

	return categories
}
