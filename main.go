package main

import (
	"go-gin/routes"
	"go-gin/store"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

func main() {

	dsn := "root:test@123!@tcp(127.0.0.1:3306)/golang_microservice"
	store.Connection(dsn)
	router := gin.Default()
	routes.UserRoutes(router)

	router.Run(":8080")
}
