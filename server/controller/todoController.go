package controller

import (
	"net/http"
	"server/model"

	"github.com/gin-gonic/gin"
)

func FetchAllTodos(c *gin.Context) {
	todos := model.FindAllTodos()

	// URLへのアクセスに対してJSONを返す
	c.JSON(http.StatusOK, todos)
}
