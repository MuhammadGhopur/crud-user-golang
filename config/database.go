package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db := "root:@tcp(127.0.0.1:3306)/golang_db?parseTime=true"

	database, err := gorm.Open(mysql.Open(db), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terhubung!")
	}

	fmt.Println("Berhasil terhubung!")

	DB = database
}
