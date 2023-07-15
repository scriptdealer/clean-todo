package storage

import "github.com/scriptdealer/clean-todo/gnosis"

type ToDoStore interface {
	GetOne(id int) gnosis.TodoItem
	GetAll() []gnosis.TodoItem
	Create(item *gnosis.TodoItem) int
	Update(item *gnosis.TodoItem) bool
	Delete(id int) bool
}

type ToDoStorage struct {
	db           map[int]gnosis.TodoItem
	currentIndex int
}

func NewMemoryStorage() *ToDoStorage {
	return &ToDoStorage{db: make(map[int]gnosis.TodoItem)}
}

func (tds *ToDoStorage) GetOne(id int) gnosis.TodoItem {
	result, found := tds.db[id]
	if found {
		return result
	}
	return result // no error
}

func (tds *ToDoStorage) GetAll() []gnosis.TodoItem {
	result := make([]gnosis.TodoItem, 0)
	for _, v := range tds.db {
		result = append(result, v)
	}
	return result
}

func (tds *ToDoStorage) Create(item *gnosis.TodoItem) int {
	tds.currentIndex++
	tds.db[tds.currentIndex] = *item
	return tds.currentIndex
}

func (tds *ToDoStorage) Update(item *gnosis.TodoItem) bool {
	_, found := tds.db[item.Id]
	if found {
		tds.db[item.Id] = *item
	}
	return found
}

func (tds *ToDoStorage) Delete(id int) bool {
	_, found := tds.db[id]
	if found {
		delete(tds.db, id)
	}
	return found
}
