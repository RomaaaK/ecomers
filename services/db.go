package services

import (
	"example/ecomers/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {

	dsn := "root:1234@tcp(localhost:3306)/ecomers?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Brand{}, &models.Category{}, &models.SubCategory{})

	var brands []models.Brand
	var category []models.Category

	db.Find(&brands)
	db.Model(&models.Category{}).Preload("Children").Find(&category)

	if len(category) == 0 {

		nike := models.SubCategory{Name: "Nike"}
		adidas := models.SubCategory{Name: "Adidas"}
		zara := models.SubCategory{Name: "Zara"}

		sportsware := models.Category{Name: "SPORTSWEAR", Children: []models.SubCategory{nike, adidas, zara}}
		mens := models.Category{Name: "Mens", Children: []models.SubCategory{nike, adidas, zara}}
		womens := models.Category{Name: "Womens", Children: []models.SubCategory{nike, adidas, zara}}

		db.Create(&sportsware)
		db.Create(&mens)
		db.Create(&womens)
	}

	if len(brands) == 0 {
		nike := models.Brand{Name: "Nike"}
		zara := models.Brand{Name: "Zara"}
		adidas := models.Brand{Name: "Adidas"}
		uniqlo := models.Brand{Name: "Uniqlo"}

		db.Create(&nike)
		db.Create(&zara)
		db.Create(&adidas)
		db.Create(&uniqlo)
	}

	DB = db
}
