package repositories

import "github.com/leandroyyy/poc-tech-challenge/src/core/domain/entities"

type ClientRepository interface {
	Create(client *entities.Client) (*entities.Client, error)
}
