package repository_impl

import (
	postgres_database "github.com/leandroyyy/poc-tech-challenge/src/adapter/driven/database/gorm"
	"github.com/leandroyyy/poc-tech-challenge/src/core/domain/entities"
)

type ClientRepositoryImpl struct {
}

func (c *ClientRepositoryImpl) Create(client *entities.Client) (*entities.Client, error) {
	postgres_database.DB.Create(&client)

	return client, nil
}
