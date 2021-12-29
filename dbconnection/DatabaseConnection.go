package dbconnection

import (
	"example/login/model"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDBConnection() {
	connectionString := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + os.Getenv("DB_OPTIONS")
	dbConnection, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Println("Connection failed")
	}
	db = dbConnection
	db.AutoMigrate(&model.User{})
}

func GetDBConnection() *gorm.DB {
	return db
}
