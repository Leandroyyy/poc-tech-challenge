package usecases

import (
	"github.com/google/uuid"
	"github.com/leandroyyy/poc-tech-challenge/src/core/application/repositories"
	"github.com/leandroyyy/poc-tech-challenge/src/core/domain/entities"
)

type CreateClientUseCase struct {
	ClientRepository repositories.ClientRepository
}

func (s *CreateClientUseCase) Execute(name string, cpf string) (*entities.Client, error) {

	client := entities.Client{
		Id: uuid.New().String(), Name: name, Cpf: cpf,
	}

	return s.ClientRepository.Create(&client)
}
