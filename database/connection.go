package database

import (
	"database/sql"

	"github.com/ankiitdev/go-auth-fiber/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	sqlDB, _ := sql.Open("mysql", "root:root123@tcp(127.0.0.1:3306)/restapigolang")
	connection, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		panic("couldn't connect with sql db")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
