package controller

import (
	"net/http"
	"server/service"

	"github.com/gin-gonic/gin"
)

func FetchAll(c *gin.Context) {
	todos := service.FindAllTodos()

	c.JSON(http.StatusOK, todos)
}
