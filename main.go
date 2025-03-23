package main

import (
	"go-todo-api/handlers"
	"go-todo-api/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// データベースの初期化
	models.InitDB()

	// ルーターの設定
	router := mux.NewRouter()

	// ルートの定義
	router.HandleFunc("/todos", handlers.GetAllTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", handlers.GetTodo).Methods("GET")
	router.HandleFunc("/todos", handlers.CreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", handlers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", handlers.DeleteTodo).Methods("DELETE")
	// 復元エンドポイントの追加
	router.HandleFunc("/todos/{id}/restore", handlers.RestoreTodo).Methods("POST")

	// 物理削除用のエンドポイント（オプション）
	router.HandleFunc("/todos/{id}/hard", handlers.HardDeleteTodo).Methods("DELETE")

	// サーバーの起動
	log.Println("サーバーを起動中: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
