package main

import (
	"fmt"
	"net/http"

	"./config"
	"./controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type appError struct {
	Message string
	Code    int
}

func (err *appError) Error() string {
	return fmt.Sprintf("[%v error] %v", err.Code, err.Message)
}

func main() {
	// ==================================== INIT ====================================

	db := config.DBInit()
	controller := &controllers.InDB{DB: db}

	// ==================================== ROUTER ====================================

	router := gin.Default()

	router.Use(errorHandlerMiddleware)

	router.GET("/person", controller.GetPersons)
	router.GET("/person/:id", controller.GetPerson)
	router.POST("/person", controller.CreatePerson)
	router.DELETE("/person/:id", controller.DeletePerson)
	router.PUT("/person/:id", controller.UpdatePerson)
	router.GET("/error", func(c *gin.Context) {
		panic(&appError{Code: http.StatusInternalServerError, Message: "hehehe error 500"})
	})

	router.Run(":8080")
}

func errorHandlerMiddleware(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			err := err.(*appError)
			c.JSON(err.Code, gin.H{
				"error":  err.Message,
				"status": err.Code,
			})
		}
	}()

	c.Next()
}
