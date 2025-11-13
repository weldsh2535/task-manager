package main

import (
	"fmt"
	"log"
	"os"

	"github.com/weldsh2535/task-manager/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = "root:@tcp(127.0.0.1:3306)/taskdb?charset=utf8mb4&parseTime=True&loc=Local"
	}

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to DB: %v", err)
	}

	fmt.Println("✅ Connected to MySQL Database")

	DB.AutoMigrate(&models.User{}, &models.Project{}, &models.Task{})
}
