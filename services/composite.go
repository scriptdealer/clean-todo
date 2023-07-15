package services

import (
	"net/http"
	"os"

	"github.com/scriptdealer/clean-todo/storage"
)

//Это т.н. usecase, а по-нашему, сервис композиции сущностей

type MainContext struct {
	// Config     *Configuration
	Db           *storage.ToDoStorage
	Server       *http.Server
	Interruption chan os.Signal
	ToDos        *todoService
	// Users        *userService
}

func NewComposer(db *storage.ToDoStorage) *MainContext {
	return &MainContext{
		Db:           db,
		ToDos:        NewToDoService(db),
		Interruption: make(chan os.Signal, 1),
	}
}
