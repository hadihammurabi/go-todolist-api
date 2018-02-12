package server

import (
	"TodoAPI/server/todo"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func Listen(host string) {
	router.HandleFunc("/", RootHandler).Methods("GET")
	todo.CreateRoute(router)
	log.Fatal(http.ListenAndServe(host, router))
}

func RootHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("/api"))
}
