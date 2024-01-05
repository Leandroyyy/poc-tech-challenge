package postgres_database

import (
	"log"

	database_model "github.com/leandroyyy/poc-tech-challenge/src/adapter/driven/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToDatabase() {
	conectionString := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err := gorm.Open(postgres.Open(conectionString))

	if err != nil {
		log.Panic("Erro ao conectar ao banco de dados")
	}

	DB.AutoMigrate(&database_model.ClientModel{})
}
