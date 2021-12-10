package main

import (
	"example/login/controller"
	"example/login/dbconnection"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbconnection.InitDBConnection()

	router := gin.Default()

	router.POST("/register", controller.AddUser)
	router.GET("/login", controller.LoginUser)

	router.Run(":8090")
}
