package todolist

import (
	"encoding/json"
	"net/http"
)

type storage interface {
  Load(key string) (interface{}, bool)
  Remove(key string)
  Exist(key string) bool
  Save(key string, value interface{}) error
  LoadAll() (map[string]interface{}, error)
}

// Todo simple todo struct
type Todo struct {
  ID string `json:"id"`
  Name string `json:"name"`
  Completed bool `json:"completed"`
}

// Handler is responsible for handling the Todolist endpoints
type Handler struct{
  storage storage
}

// NewHandler is the constructor for the todolist.Handler
func NewHandler(storage storage) *Handler {
  return &Handler{
    storage: storage,
  }
}

// Create handles the creation of new todos
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
  todo := Todo{}

  err := json.NewDecoder(r.Body).Decode(&todo)
  if err != nil {
    JSONError(w, err, http.StatusBadRequest)
		return 
  }

  err = h.storage.Save(todo.ID, todo)
  if err != nil {
    JSONError(w, err, http.StatusInternalServerError)
    return
  }

  JSONSuccess(w, todo, http.StatusCreated)
}

// JSONSuccess is a JSON success response.
func JSONSuccess(w http.ResponseWriter, response interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

// JSONError is a JSON error response.
func JSONError(w http.ResponseWriter, err error, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status_code": statusCode,
		"message":     err.Error(),
	})
}
