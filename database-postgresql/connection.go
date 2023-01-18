package database_postgresql

import (
	"log"
	"user-api-rest/database-postgresql/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var Connection_db *gorm.DB

func Connect(database_config string) {
	var connection_error error
	Connection_db, connection_error = gorm.Open(postgres.Open(database_config), &gorm.Config{})

	if connection_error != nil {
		log.Fatal(connection_error)
		panic(connection_error)
	}

	log.Println("Database POSTGRESQL connection sucessfully")
}

func CreateDatabase() {
	if Connection_db != nil {
		Connection_db.AutoMigrate(model.User{})
	}
}