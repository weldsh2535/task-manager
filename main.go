package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Load environment variables (optional)
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		// Example DSN: username:password@tcp(127.0.0.1:3306)/taskdb?charset=utf8mb4&parseTime=True&loc=Local
		dsn = "root:@tcp(127.0.0.1:3306)/taskdb?charset=utf8mb4&parseTime=True&loc=Local"
	}

	// Connect to MySQL
	db, err := InitDB(dsn)
	if err != nil {
		log.Fatalf("❌ Failed to connect to DB: %v", err)
	}

	// Migrate schema
	if err := db.AutoMigrate(&Task{}); err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	// Create router
	r := NewRouter(db)

	addr := ":8080"
	log.Printf("🚀 Server running on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
