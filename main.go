package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vitorf7/todo_go/internal/storage"
	"github.com/vitorf7/todo_go/internal/todolist"
)

func main() {
	router, err := run()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Starting the server on port 8080")
	if err := http.ListenAndServe(":8080", Middleware(router)); err != nil {
		log.Fatal(err)
	}
}

func run() (*mux.Router, error) {
  memoryStorage := storage.NewStore()
	todoListHandler := todolist.NewHandler(memoryStorage)

	r := mux.NewRouter()
	r.HandleFunc("/todo", todoListHandler.Create).Methods(http.MethodPost)
	return r, nil
}

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		h.ServeHTTP(w, r)
	})
}
