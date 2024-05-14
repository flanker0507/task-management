package configs

import (
	"fmt"
	"go-todo-app/database"
	"go-todo-app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var db_name = ""
var db_port = "3306"
var db_user = "root"
var db_password = "root"
var db_host = "127.0.0.1"

func bootDatabase() {
	if dbNameEnv := os.Getenv("DB_NAME"); dbNameEnv != "" {
		db_name = dbNameEnv
	}

	if dbPortEnv := os.Getenv("DB_PORT"); dbPortEnv != "" {
		db_port = dbPortEnv
	}

	if dbUserEnv := os.Getenv("DB_USER"); dbUserEnv != "" {
		db_user = dbUserEnv
	}

	if dbPasswordEnv := os.Getenv("DB_PASSWORD"); dbPasswordEnv != "" {
		db_password = dbPasswordEnv
	}

	if dbHostEnv := os.Getenv("DB_HOST"); dbHostEnv != "" {
		db_host = dbHostEnv
	}
}

func connectDatabase() {
	var errConnection error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_password, db_host, db_port, db_name)
	database.DB, errConnection = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if errConnection != nil {
		panic("Can't connect to database")
	} else {
		log.Println("Connected to database")
	}
}

func runMigration() {
	err := database.DB.AutoMigrate(
		models.Todo{},
		models.User{},
	)
	if err != nil {
		fmt.Println(err)
		log.Println("Failed to Migrate schema")
	} else {
		log.Println("schemas migrated")
	}
}
