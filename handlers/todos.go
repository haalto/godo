package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/haalto/godo/services"
)

type Todos struct {
	logger *log.Logger
}

func NewTodos(logger *log.Logger) *Todos {
	return &Todos{logger}
}

func (t *Todos) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t.getTodos(w, r)
	}
}

func (t *Todos) getTodos(w http.ResponseWriter, r *http.Request) {
	t.logger.Print("Getting todos")
	todos := services.GetTodos()
	encoder := json.NewEncoder(w)
	err := encoder.Encode(todos)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
