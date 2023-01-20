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

	db.AutoMigrate(&models.Brand{})

	var brands []models.Brand

	db.Find(&brands)

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
