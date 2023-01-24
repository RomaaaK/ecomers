package services

import (
	"errors"
	"example/ecomers/models"
)

func GetProducts() []models.Product {
	var products []models.Product

	DB.Find(&products)

	return products
}

func GetProductById(id int) (models.Product, error) {
	var product models.Product

	DB.Preload("Brand").First(&product, id)

	if product.ID == 0 {
		return models.Product{}, errors.New("product not found")
	}

	return product, nil
}

func GetProductsByBrand(id int) []models.Product {
	var products []models.Product

	DB.Where("brand_id = ?", id).Find(&products)

	return products
}

func GetProductsByCategory(id int) []models.Product {
	var products []models.Product

	DB.Where("category_id = ?", id).Find(&products)

	return products
}
