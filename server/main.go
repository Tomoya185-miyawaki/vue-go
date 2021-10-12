package main

import (
	"log"
	"server/controller"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/products", controller.FetchAllProducts)

	err := router.Run(":81")
	if err != nil {
		log.Fatal("Server Error: ", err)
	}
}
