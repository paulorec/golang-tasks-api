package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Task struct {
	Id   string `json:"id,omitempty"`
	Text string `json:"text,omitempty"`
	From *User  `json:"from,omitempty"`
	To   *User  `json:"to,omitempty"`
}

type User struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

var tasks []Task

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/task", GetTasks).Methods("GET")
	router.HandleFunc("/task/{id}", GetTask).Methods("GET")
	router.HandleFunc("/task/{id}", CreateTask).Methods("POST")
	router.HandleFunc("/task/{id}", DeleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))

	addSampleData()
}

func addSampleData() {

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
