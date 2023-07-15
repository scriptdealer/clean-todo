package services

import (
	"context"

	"github.com/scriptdealer/clean-todo/gnosis"
	"github.com/scriptdealer/clean-todo/storage"
)

type todoService struct {
	store storage.ToDoStore
}

func NewToDoService(db storage.ToDoStore) *todoService {
	return &todoService{store: db}
}

func (tds *todoService) Create(title, description string) int {
	item := gnosis.TodoItem{
		Title:       title,
		Description: description,
	}
	return tds.store.Create(&item)
}

func (tds *todoService) Update(id int, title, description string, done bool) bool {
	patch := gnosis.TodoItem{
		Id:          id,
		Title:       title,
		Description: description,
		Done:        done,
	}
	return tds.store.Update(&patch)
}

func (tds *todoService) Delete(id int) bool {
	return tds.store.Delete(id)
}

func (tds *todoService) Get(id int) gnosis.TodoItem {
	return tds.store.GetOne(id)
}

func (tds *todoService) GetAll(ctx context.Context) []gnosis.TodoItem {
	return tds.store.GetAll()
}
