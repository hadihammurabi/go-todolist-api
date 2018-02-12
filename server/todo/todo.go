package todo

import (
	"TodoAPI/database"
	"TodoAPI/structs"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// CreateRoute - Create route for Todo
func CreateRoute(router *mux.Router) {
	router.HandleFunc("/todo", todoHandler).Methods("GET")
	router.HandleFunc("/todo/{id}", todoHandler).Methods("GET")
	router.HandleFunc("/todo", todoInsertHandler).Methods("POST")
	router.HandleFunc("/todo/{id}", todoHandler).Methods("PUT")
	router.HandleFunc("/todo/{id}", todoDeleteHandler).Methods("DELETE")
}

// TodosHandler - Handle /todo
func todoHandler(w http.ResponseWriter, req *http.Request) {
	var todos []structs.Todo
	database.GetTodos().All(&todos)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func todoInsertHandler(w http.ResponseWriter, req *http.Request) {
	var todo structs.Todo
	var message *structs.Message
	json.NewDecoder(req.Body).Decode(&todo)
	err := database.AddTodo(todo.Title)
	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		message = &structs.Message{Code: http.StatusInternalServerError, Title: "Error", Body: "Error when save your new todo item"}
		panic(err)
	} else {
		w.WriteHeader(http.StatusCreated)
		message = &structs.Message{Code: http.StatusCreated, Title: "New Todo Item Added", Body: "Your new todo has been saved"}
	}
	json.NewEncoder(w).Encode(message)
}

func todoDeleteHandler(w http.ResponseWriter, req *http.Request) {
	var message *structs.Message
	id := mux.Vars(req)["id"]
	err := database.DeleteTodo(id)
	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		message = &structs.Message{Code: http.StatusInternalServerError, Title: "Error", Body: "Error when delete your todo item"}
		fmt.Println(err)
	} else {
		w.WriteHeader(http.StatusCreated)
		message = &structs.Message{Code: http.StatusOK, Title: "A Todo Item Deleted", Body: "Your todo has been deleted"}
	}
	json.NewEncoder(w).Encode(message)
}
