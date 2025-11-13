package handlers

import "gorm.io/gorm"

// DB is the package-level database handle used by handler functions.
var DB *gorm.DB

// SetDB assigns the DB handle for the handlers package.
func SetDB(db *gorm.DB) {
	DB = db
}
