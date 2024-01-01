package repository_impl

import (
	"github.com/leandroyyy/poc-tech-challenge/src/core/domain/entities"
	postgres_database "github.com/leandroyyy/poc-tech-challenge/src/infra/database/gorm"
)

type ClientRepositoryImpl struct {
}

func (c *ClientRepositoryImpl) Create(client *entities.Client) (*entities.Client, error) {
	postgres_database.DB.Create(&client)

	return client, nil
}
