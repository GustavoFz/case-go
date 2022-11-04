package config

import (
	"fmt"
	"log"

	"github.com/gustavofz/case-eulabs/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var envs map[string]string = InitEnv()

var (
	DBHost = envs["DBHost"]
	DBPort = envs["DBPort"]
	DBUser = envs["DBUser"]
	DBPass = envs["DBPass"]
	DBName = envs["DBName"]
)

func getDBConnection() string {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DBUser,
		DBPass,
		DBHost,
		DBPort,
		DBName,
	)
	return connection
}

func InitConnection() *gorm.DB {
	dsn := getDBConnection()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("Error connection")
	}

	// Create product table
	db.AutoMigrate(&model.Product{})

	return db
}
