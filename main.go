package main

import (
	"go-crud-restapi/controllers/authcontroller"
	"go-crud-restapi/controllers/productcontroller"
	"go-crud-restapi/initializers"
	"go-crud-restapi/middlewares"
	"go-crud-restapi/models"
	"os"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
}

func main() {
	port := os.Getenv("PORT")
	r := gin.Default();
	models.ConnectDatabase()

	r.POST("/api/login", authcontroller.Login)
	r.POST("/api/register",authcontroller.Register)
	r.GET("/api/logout",authcontroller.Logout)

	r.GET("/api/product",middlewares.JWTMiddleware, productcontroller.Index)
	r.GET("/api/product/:id",productcontroller.Show)
	r.POST("/api/product",productcontroller.Create)
	r.PUT("/api/product/:id",productcontroller.Update)
	r.DELETE("/api/product",productcontroller.Delete)

	r.Run(":"+port)
}