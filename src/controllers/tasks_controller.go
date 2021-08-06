package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/janper231/learning-golang/db"
	"github.com/janper231/learning-golang/src/models"
)

func GetTasks(c *gin.Context) {
	response := []*models.Task{}

	db.DB(&response)

	c.JSON(200, gin.H{"status": true, "data": response})
}

func CreateTask(c *gin.Context) {
	response := []*models.Task{}
	body := &models.Task{}

	c.ShouldBindJSON(&body)

	db.DB(&response)

	body.ID = len(response) + 1

	c.JSON(200, gin.H{"status": "true", "data": body})
}

func GetTask(c *gin.Context) {
	Id := c.Param("id")
	response := []*models.Task{}

	db.DB(&response)

	for _, task := range response {
		id, _ := strconv.Atoi(Id)

		if task.ID == id {
			c.JSON(200, gin.H{"status": "true", "data": task})
		}
	}
}

func DeleteTask(c *gin.Context) {
	taskId := c.Param("id")
	response := []*models.Task{}

	db.DB(&response)

	for index, task := range response {
		id, _ := strconv.Atoi(taskId)

		if task.ID == id {
			response = append(response[:index], response[index+1:]...)

			c.JSON(200, gin.H{"status": "true", "data": response})
		}
	}
}

func UpdateTask(c *gin.Context) {
	IdParam := c.Param("id")
	body := &models.Task{}
	tasks := []*models.Task{}

	c.ShouldBindJSON(&body)

	db.DB(&tasks)

	for index, task := range tasks {
		Id, _ := strconv.Atoi(IdParam)

		if task.ID == Id {
			tasks = append(tasks[:index], tasks[index+1:]...)
			body.ID = task.ID
			tasks = append(tasks, body)

			c.JSON(200, gin.H{"status": "true", "data": tasks})
			break
		}
	}
}
