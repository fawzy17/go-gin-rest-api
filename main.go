package main

import (
	"github.com/fawzy17/rest-api-gin-gorm/config"
	"github.com/fawzy17/rest-api-gin-gorm/controllers/productcontroller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDatabase()

	r.GET("/api/v1/products", productcontroller.Index)
	r.GET("/api/v1/products/:id", productcontroller.Show)
	r.POST("/api/v1/products", productcontroller.Create)
	r.PUT("/api/v1/products/:id", productcontroller.Update)
	r.DELETE("/api/v1/products/:id", productcontroller.Delete)

	r.Run()
}
