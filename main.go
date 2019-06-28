package main

import (
	"./config"
	"./controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// ==================================== INIT ====================================

	db := config.DBInit()
	controller := &controllers.InDB{DB: db}

	// ==================================== ROUTER ====================================

	router := gin.Default()

	router.GET("/person", controller.GetPersons)
	router.GET("/person/:id", controller.GetPerson)
	router.POST("/person", controller.CreatePerson)
	router.DELETE("/person/:id", controller.DeletePerson)
	router.PUT("/person/:id", controller.UpdatePerson)

	router.Run(":8080")
}
