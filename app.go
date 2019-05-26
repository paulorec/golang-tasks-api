package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) start(addr string) {
	a.Router = mux.NewRouter()
	a.addSampleData()
	a.initRoutes()
	log.Fatal(http.ListenAndServe(addr, a.Router))

}

func (app *App) initRoutes() {
	app.Router.HandleFunc("/task", GetTasks).Methods("GET")
	app.Router.HandleFunc("/task/{id}", GetTask).Methods("GET")
	app.Router.HandleFunc("/task/{id}", CreateTask).Methods("POST")
	app.Router.HandleFunc("/task/{id}", DeleteTask).Methods("DELETE")
}

var tasks []Task

func (app *App) addSampleData() {

	var lebowski = &User{Id: "1", Name: "Lebowski", Email: "the@dude.com"}

	tasks = append(tasks, Task{Id: "1", Text: "In a rocks glass combine vodka and coffee liqueur over ice", From: lebowski, To: lebowski})
	tasks = append(tasks, Task{Id: "1", Text: "Add milk or cream and stir", From: lebowski, To: lebowski})

}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range tasks {
		if item.Id == params["id"] {

			json.NewEncoder(w).Encode(item)
			return
		}

	}
	json.NewEncoder(w).Encode(&Task{})
}
func CreateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.Id = params["id"]

}
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range tasks {
		if item.Id == params["id"] {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(tasks)
}
