package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Task struct {
	ID      int    `json:ID`
	Name    string `json:Name`
	Content string `json:Content`
}

type allTasks []Task

var tasks = allTasks{
	{
		ID:      1,
		Name:    "tarea uno",
		Content: "es la tarea uno",
	},
}

func getTasks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, "Insert valid data")
	}

	json.Unmarshal(reqBody, &newTask)

	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to api test")
}

func getTask(w http.ResponseWriter, r *http.Request) {
	taskId := getParams(w, r, "id")

	for _, task := range tasks {
		if task.ID == taskId {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
		}
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	taskId := getParams(w, r, "id")

	for index, task := range tasks {
		if task.ID == taskId {
			tasks = append(tasks[:index], tasks[index+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(tasks)
		}
	}
}

func updateTask(w http.ResponseWriter, r *http.Request) {
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

func getParams(w http.ResponseWriter, r *http.Request, field string) int {
	vars := mux.Vars(r)
	value, err := strconv.Atoi(vars[field])

	if err != nil {
		fmt.Fprintf(w, "Invalid Id")
	}

	return value
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/task/{id}", updateTask).Methods("PUT")
	router.HandleFunc("/", indexRoute)

	log.Fatal(http.ListenAndServe(":8080", router))
}
