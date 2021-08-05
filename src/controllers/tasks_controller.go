package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/janper231/learning-golang/db"
	"github.com/janper231/learning-golang/src/models"
)

func GetTasks(c *gin.Context) {
	c.JSON(200, gin.H{"status": "posted"})
}

func CreateTask(c *gin.Context) {
	var newTask models.Task
	reqBody, err := ioutil.ReadAll(c.Request.Body)
	tasks := db.DB()

	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"status": "false", "message": "Insert valid data"})
	}

	json.Unmarshal(reqBody, &newTask)

	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	c.JSON(http.StatusAccepted, gin.H{"status": "true", "data": newTask})
}

func GetTask(c *gin.Context) {
	taskId := getParams(w, r, "id")

	for _, task := range tasks {
		if task.ID == taskId {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
		}
	}
}

func DeleteTask(c *gin.Context) {
	taskId := getParams(w, r, "id")

	for index, task := range tasks {
		if task.ID == taskId {
			tasks = append(tasks[:index], tasks[index+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(tasks)
		}
	}
}

func UpdateTask(c *gin.Context) {
	taskId := getParams(w, r, "id")
	var updateTask Task

	reqBoyd, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Invalid Id")
	}

	json.Unmarshal(reqBoyd, &updateTask)

	for index, task := range tasks {
		if task.ID == taskId {
			tasks = append(tasks[:index], tasks[index+1:]...)
			updateTask.ID = taskId
			tasks = append(tasks, updateTask)

			json.NewEncoder(w).Encode(updateTask)
		}
	}
}
