package database

import (
	"go-checkout/src/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	DB, err = gorm.Open(mysql.Open("root:root@tcp(checkout_db:3306)/checkout"), &gorm.Config{})

	if err != nil {
		panic("Could not connect with the database!" + err.Error())
	}
}

func AutoMigrate() {
	DB.AutoMigrate(models.Link{},models.Order{},models.OrderItem{},models.Product{},models.KafkaError{})
}
