package main

import (
	routes "github.com/leandroyyy/poc-tech-challenge/src/cmd/api"
	postgres_database "github.com/leandroyyy/poc-tech-challenge/src/infra/database/gorm"
)

func main() {
	postgres_database.ConnectToDatabase()
	routes.HandlerRequests()
}
