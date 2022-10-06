package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/vitorf7/todo_go/internal/storage"
	"github.com/vitorf7/todo_go/internal/todolist"
	"gotest.tools/assert"
)

func Test_CreateTodoItem(t *testing.T) {
  memoryStorage := storage.NewStore()
	todoListHandler := todolist.NewHandler(memoryStorage)


  todoRequest := `{"id": "4dfe170b-d12c-4d75-aca6-6dde870e7a44", "name": "My todo", "completed": true}`

  writer := httptest.NewRecorder()
  request := httptest.NewRequest(http.MethodPost, "/todo", strings.NewReader(todoRequest))
  
  todoListHandler.Create(writer, request)

  todo, ok := memoryStorage.Load("4dfe170b-d12c-4d75-aca6-6dde870e7a44")
  if !ok {
    t.Fail()
  }
  actualTodo, ok := todo.(todolist.Todo)
    
  assert.Equal(t, http.StatusCreated, writer.Code)
  assert.Equal(t, "4dfe170b-d12c-4d75-aca6-6dde870e7a44", actualTodo.ID)
}
