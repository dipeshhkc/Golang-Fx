package infrastructure

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

// Database modal
type Database struct {
	DB *gorm.DB
}

// NewDatabase creates a new database instance
func NewDatabase(logger Logger) Database {

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)

	db, err := gorm.Open("mysql", url)

	if err != nil {
		logger.Zap.Info("Url: ", url)
		logger.Zap.Panic(err)
	}

	logger.Zap.Info("Database connection established")

	return Database{
		DB: db,
	}
}
