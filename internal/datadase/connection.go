package datadase

import (
	"awesomeProject1/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	connection, err := gorm.Open(mysql.Open("root:12345678@/adp"), &gorm.Config{})
	if err != nil {
		panic("could not connect to database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
