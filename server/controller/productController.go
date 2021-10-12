package controller

import (
	"net/http"
	"server/model"

	"github.com/gin-gonic/gin"
)

func FetchAllProducts(c *gin.Context) {
	resultProducts := model.FindAllProducts()

	// URLへのアクセスに対してJSONを返す
	c.JSON(http.StatusOK, resultProducts)
}
