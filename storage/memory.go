package storage

import (
	"sync"

	"github.com/scriptdealer/clean-todo/gnosis"
)

type ToDoStore interface {
	GetOne(id int) gnosis.TodoItem
	GetAll() []gnosis.TodoItem
	Create(item *gnosis.TodoItem) int
	Update(item *gnosis.TodoItem) bool
	Delete(id int) bool
}

type ToDoStorage struct {
	db           map[int]gnosis.TodoItem
	lock         sync.Mutex
	currentIndex int
}

func NewMemoryStorage() *ToDoStorage {
	return &ToDoStorage{db: make(map[int]gnosis.TodoItem)}
}

func (tds *ToDoStorage) GetOne(id int) gnosis.TodoItem {
	tds.lock.Lock()
	defer tds.lock.Unlock()
	result, found := tds.db[id]
	if found {
		return result
	}
	return result // no error
}

func (tds *ToDoStorage) GetAll() []gnosis.TodoItem {
	tds.lock.Lock()
	defer tds.lock.Unlock()
	result := make([]gnosis.TodoItem, 0)
	for _, v := range tds.db {
		result = append(result, v)
	}
	return result
}

func (tds *ToDoStorage) Create(item *gnosis.TodoItem) int {
	tds.lock.Lock()
	defer tds.lock.Unlock()
	tds.currentIndex++
	tds.db[tds.currentIndex] = *item
	return tds.currentIndex
}

func (tds *ToDoStorage) Update(item *gnosis.TodoItem) bool {
	tds.lock.Lock()
	defer tds.lock.Unlock()
	_, found := tds.db[item.Id]
	if found {
		tds.db[item.Id] = *item
	}
	return found
}

func (tds *ToDoStorage) Delete(id int) bool {
	tds.lock.Lock()
	defer tds.lock.Unlock()
	_, found := tds.db[id]
	if found {
		delete(tds.db, id)
	}
	return found
}
