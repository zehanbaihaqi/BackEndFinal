package configs

import (
	"sim/models/pemberitahuan"
	"sim/models/users"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	conn := "root:zehan123@tcp(127.0.0.1:3306)/sim?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(conn), &gorm.Config{})

	if err != nil {
		panic("Error to connect to the database")
	}
	Migration()
}

func Migration() {
	DB.AutoMigrate(users.Users{}, pemberitahuan.Pemberitahuans{})
}
