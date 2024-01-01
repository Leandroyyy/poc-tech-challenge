package driver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	repository_impl "github.com/leandroyyy/poc-tech-challenge/src/adapter/driven/database"
	driver_dto "github.com/leandroyyy/poc-tech-challenge/src/adapter/driver/dtos"
	usecases "github.com/leandroyyy/poc-tech-challenge/src/core/application/use-cases"
)

func CreateClienteController(g *gin.Context) {
	var clientDto driver_dto.CreateClientDto
	if err := g.ShouldBindJSON(&clientDto); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	clientRepository := &repository_impl.ClientRepositoryImpl{}

	createClienteUseCase := usecases.CreateClientUseCase{
		ClientRepository: clientRepository,
	}

	clientCreated, err := createClienteUseCase.Execute(
		clientDto.Name,
		clientDto.Cpf,
	)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	g.JSON(http.StatusOK, clientCreated)
}
