package main

import (
	"log"
	database_postgresql "user-api-rest/database-postgresql"
	user_server "user-api-rest/server"
)

var database_config = "host=localhost port=5432 dbname=users user=jfjara password=password"
var server_port = ":8080"

func main() {
	database_postgresql.Connect(database_config)
	database_postgresql.CreateDatabase()
	
	log.Println("Starting Server at 8080")
	user_server.Start(server_port)
}
