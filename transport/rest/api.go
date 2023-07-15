package rest

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gorilla/mux"
	"github.com/scriptdealer/clean-todo/services"
)

var serviceComposer *services.MainContext

func InitHandlers(composer *services.MainContext) http.Handler {
	serviceComposer = composer
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/todo", AllItems).Methods(http.MethodGet)
	r.HandleFunc("/todo", AddItem).Methods(http.MethodPost)
	r.HandleFunc("/todo/{id}", GetItem).Methods(http.MethodGet)
	r.HandleFunc("/todo/{id}", UpdateItem).Methods(http.MethodPatch)
	r.HandleFunc("/todo/{id}", DeleteItem).Methods(http.MethodDelete)
	return r
}

func LogRecover() {
	if err := recover(); err != nil {
		fmt.Println("Fail:", string(debug.Stack()))
	}
}
