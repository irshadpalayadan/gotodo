package main

import (
	"github.com/gin-gonic/gin"
	"github.com/irshadpalayadan/gotodo/package/middleware"
	"github.com/irshadpalayadan/gotodo/package/todo"

	log "github.com/sirupsen/logrus"
)

func getServerStatus(ctx *gin.Context) {

	ctx.JSON(200, "Serever running successfully")
}

func main() {
	router := gin.Default()
	router.Use(middleware.GlobalMiddleware())

	router.GET("/status", getServerStatus)
	router.GET("/todos", todo.GetTodo)
	router.POST("/todos", todo.AddTodo)
	router.Run(":8000")

	defer log.Fatal("Router creation failed")

}
