package database

import (
	"TodoAPI/structs"
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var col = GetConnection().DB("todoList").C("todos")

// GetTodos - Return all todos
func GetTodos() *mgo.Query {
	return col.Find(nil)
}

// AddTodo - Create a Todo item
func AddTodo(title string) error {
	tmp := &structs.Todo{Title: title}
	return col.Insert(tmp)
}

// DeleteTodo - Delete a Todo item
func DeleteTodo(id string) error {
	fmt.Println(bson.M{"_id": id})
	return col.Remove(bson.M{"_id": id})
}
