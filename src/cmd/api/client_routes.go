package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/leandroyyy/poc-tech-challenge/src/adapter/driver"
)

func HandlerRequests() {
	r := gin.Default()

	r.POST("/alunos", driver.CreateClienteController)

	r.Run()
}
