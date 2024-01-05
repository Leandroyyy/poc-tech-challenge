package main

import (
	postgres_database "github.com/leandroyyy/poc-tech-challenge/src/adapter/driven/database/gorm"
	routes "github.com/leandroyyy/poc-tech-challenge/src/adapter/driver/http"
)

func main() {
	postgres_database.ConnectToDatabase()
	routes.HandlerRequests()
}
