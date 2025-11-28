package main

import (
	"log"
	"net/http"

	httpinternal "example.com/todo-rest/internal/http"
	"example.com/todo-rest/internal/todo"
)

func main() {
	store := todo.NewStore("data/todos.json")
	h := httpinternal.NewHandler(store)

	s := &http.Server{
		Addr: ":8080",
		Handler: h,
	}

	log.Fatal(s.ListenAndServe())
}
