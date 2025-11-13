// package main

// import (
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// // InitDB opens a connection to the database using the provided DSN
// // and returns the gorm DB handle (and any error) so callers can use it.
// func InitDB(dsn string) (*gorm.DB, error) {
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	DB = db
// 	return db, nil
// }

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
