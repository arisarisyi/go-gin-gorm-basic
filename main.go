package main

import (
	"go-crud-restapi/controllers/productcontroller"
	"go-crud-restapi/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default();
	models.ConnectDatabase()

	r.GET("/api/product",productcontroller.Index)
	r.GET("/api/product/:id",productcontroller.Show)
	r.POST("/api/product",productcontroller.Create)
	r.PUT("/api/product/:id",productcontroller.Update)
	r.DELETE("/api/product",productcontroller.Delete)

	r.Run(":3030")
}