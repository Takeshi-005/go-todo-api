package handlers

import (
	"encoding/json"
	"go-todo-api/models"
	"net/http"

	"github.com/gorilla/mux"
)

// TodoRequest はTodoのリクエストボディを表す
type TodoRequest struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// GetAllTodos 全てのTodoを取得するハンドラー
func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos := models.GetAllTodos()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// GetTodo 指定されたIDのTodoを取得するハンドラー
func GetTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	todo, err := models.GetTodoByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Todo not found"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

// CreateTodo 新しいTodoを作成するハンドラー
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var req TodoRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	if req.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Title is required"})
		return
	}

	newTodo := models.CreateTodo(req.Title)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}

// UpdateTodo 既存のTodoを更新するハンドラー
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var req TodoRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	updatedTodo, err := models.UpdateTodo(id, req.Title, req.Completed)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Todo not found"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTodo)
}

// DeleteTodo 指定されたIDのTodoを削除するハンドラー（論理削除）
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	err := models.DeleteTodo(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Todo not found or already deleted"})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// RestoreTodo 論理削除されたTodoを復元するハンドラー
func RestoreTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	restoredTodo, err := models.RestoreTodo(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(restoredTodo)
}

// HardDeleteTodo 指定されたIDのTodoを物理削除するハンドラー
func HardDeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	err := models.HardDeleteTodo(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Todo not found"})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
