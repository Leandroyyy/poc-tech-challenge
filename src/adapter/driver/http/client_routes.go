package routes

import (
	"github.com/gin-gonic/gin"
	driver "github.com/leandroyyy/poc-tech-challenge/src/adapter/driver/controllers"
)

func HandlerRequests() {
	r := gin.Default()

	r.POST("/alunos", driver.CreateClienteController)

	r.Run()
}
