package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	username = "prasetya"
	password = "1234"
	host     = "localhost"
	port     = "3306"
	dbName   = "go_crud"
)

func ConnectDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", username, password, host, port, dbName)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var models = []interface{}{&User{}, &Student{}}

	database.AutoMigrate(models...)

	DB = database
}
