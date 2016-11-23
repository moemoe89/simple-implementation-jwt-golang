package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"simple-implementation-jwt-golang/controllers"
	"time"
)

func main() {

	router := gin.Default()

	router.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type",
		ExposedHeaders: "",
		MaxAge: 50 * time.Second,
		Credentials: true,
		ValidateHeaders: false,
	}))

	router.GET("/generate",controllers.GenerateToken)
	router.GET("/parse",controllers.CheckToken)
	router.Run(":3000")

}