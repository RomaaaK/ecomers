package services

import (
	"example/ecomers/models"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDataBase() {
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")

	fmt.Println(db_user, db_pass, "-----------------------------------")

	dsn := db_user + ":" + db_pass + "@tcp(localhost:3306)/ecomers?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database")
	}

	db.AutoMigrate(&models.Brand{}, &models.Category{})

	var brands []models.Brand
	var category []models.Category

	db.Find(&brands)
	db.Model(&models.Category{}).Preload("Childrens").Find(&category)

	if len(category) == 0 {

		nike := models.Category{Name: "Nike"}
		adidas := models.Category{Name: "Adidas"}
		zara := models.Category{Name: "Zara"}

		sportsware := models.Category{Name: "SPORTSWEAR", Childrens: []models.Category{nike, adidas, zara}}
		mens := models.Category{Name: "Mens", Childrens: []models.Category{zara, adidas}}
		womens := models.Category{Name: "Womens", Childrens: []models.Category{adidas, nike}}
		other := models.Category{Name: "Other"}

		db.Create(&sportsware)
		db.Create(&mens)
		db.Create(&womens)
		db.Create(&other)
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
