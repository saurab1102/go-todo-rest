package http

import (
	//"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"example.com/todo-rest/internal/todo"
)

type Handler struct {
	store *todo.Store
}

func NewHandler(s *todo.Store) *Handler{
	return &Handler{store: s}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	p:= strings.Split(strings.Trim(r.URL.Path,"/"),"/")

	if len(p) == 1 && p[0]== "todos" {
		switch r.Method {
			case http.MethodGet:
				h.list(w,r)
			case http.MethodPost:
				h.create(w,r)
		}

		return
	}

	if len(p) == 2 && p[0]=="todos" {
		id,_:=strconv.Atoi(p[1])
		switch r.Method {
			case http.MethodDelete :
				h.delete(id)
			case http.MethodPut:
				h.done(id)
		}

		return
	}

	w.WriteHeader(http.StatusNotFound)
}


func (h *Handler) list(w http.ResponseWriter, r *http.Request){}
func (h *Handler) create(w http.ResponseWriter, r *http.Request){}
func (h *Handler) delete(id int){}
func (h *Handler) done(id int){}
