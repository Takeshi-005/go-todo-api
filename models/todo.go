package models

import (
	"errors"
	"sync"
	"time"
)

type Todo struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Completed bool       `json:"completed"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// インメモリデータベースの代わり
var (
	todos  = make(map[string]Todo)
	mutex  = &sync.Mutex{}
	nextID = 1
)

// InitDB データベースの初期化
func InitDB() {
	// 実際のアプリケーションではここでDB接続などの処理を行う
	mutex.Lock()
	defer mutex.Unlock()

	// ダミーデータの作成
	todos["1"] = Todo{
		ID:        "1",
		Title:     "Go to the gym",
		Completed: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}
	nextID = 2
}

func GetAllTodos() []Todo {
	mutex.Lock()
	defer mutex.Unlock()

	allTodos := make([]Todo, 0, len(todos))
	for _, todo := range todos {

		if todo.DeletedAt == nil {
			allTodos = append(allTodos, todo)
		}
	}

	return allTodos
}

func GetTodoByID(id string) (Todo, error) {
	mutex.Lock()
	defer mutex.Unlock()

	todo, ok := todos[id]
	if !ok {
		return Todo{}, errors.New("todo not found")
	}

	if todo.DeletedAt != nil {
		return Todo{}, errors.New("todo not found")
	}

	return todo, nil
}

func CreateTodo(title string) Todo {
	mutex.Lock()
	defer mutex.Unlock()

	id := string(rune(nextID + '0'))
	nextID++

	todo := Todo{
		ID:        id,
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}

	todos[id] = todo
	return todo
}

// UpdateTodo 既存のTODOを更新する
func UpdateTodo(id string, title string, completed bool) (Todo, error) {
	mutex.Lock()
	defer mutex.Unlock()

	todo, ok := todos[id]
	if !ok {
		return Todo{}, errors.New("todo not found")
	}

	if todo.DeletedAt != nil {
		return Todo{}, errors.New("todo not found")
	}

	todo.Title = title
	todo.Completed = completed
	todo.UpdatedAt = time.Now()
	todos[id] = todo

	return todo, nil
}

func DeleteTodo(id string) error {
	mutex.Lock()
	defer mutex.Unlock()

	todo, ok := todos[id]
	if !ok {
		return errors.New("todo not found")
	}

	if todo.DeletedAt != nil {
		return errors.New("todo not found")
	}

	now := time.Now()
	todo.DeletedAt = &now
	todos[id] = todo

	return nil
}

// HardDeleteTodo 指定されたIDのTodoを物理削除（必要に応じて）
func HardDeleteTodo(id string) error {
	mutex.Lock()
	defer mutex.Unlock()

	_, exists := todos[id]
	if !exists {
		return errors.New("todo not found")
	}

	delete(todos, id)
	return nil
}

// RestoreTodo 論理削除されたTodoを復元
func RestoreTodo(id string) (Todo, error) {
	mutex.Lock()
	defer mutex.Unlock()

	todo, exists := todos[id]
	if !exists {
		return Todo{}, errors.New("todo not found")
	}

	// 論理削除されていない場合
	if todo.DeletedAt == nil {
		return Todo{}, errors.New("todo is not deleted")
	}

	// 論理削除を解除
	todo.DeletedAt = nil
	todo.UpdatedAt = time.Now()
	todos[id] = todo

	return todo, nil
}
