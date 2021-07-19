package main

import (
	"fmt"
	"log"
	"net/http"

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

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "bienvenido a la api rest")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)

	log.Fatal(http.ListenAndServe(":8080", router))
}
