package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/janper231/learning-golang/src/controllers"
)

func main() {
	/* get port of run server */
	port := os.Getenv("PORT")

	/* init resolver routes */
	router := gin.Default()

	/* set routes */
	router.GET("/", controllers.Init)
	router.GET("/tasks", controllers.GetTasks)
	router.POST("/tasks", controllers.CreateTask)
	router.GET("/task/:id", controllers.GetTask)
	router.DELETE("/task/:id", controllers.DeleteTask)
	router.PUT("/task/:id", controllers.UpdateTask)

	router.Run(":" + port)
}
