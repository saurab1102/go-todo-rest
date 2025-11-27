package main

import (
	"log"
	"net/http"

	"example.com/todo-rest/internal/http"
	"example.com/todo-rest/internal/todo"
)

func main() {
	store := todo.NewStore("data/todos.json")
	h := http.NewHandler(store)

	s := &http.Server{
		Addr: ":8080",
		Handler: h,
	}

	log.Fatal(s.ListenAndServe())
}
