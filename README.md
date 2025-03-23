# Go TODO API

シンプルなToDo管理のためのREST APIです。Goで実装されており、タスクの作成、読み取り、更新、削除（CRUD操作）ができます。論理削除（ソフトデリート）機能も実装されています。

## 機能

- ToDo一覧の取得
- 単一のToDoの取得
- 新しいToDoの作成
- 既存のToDoの更新
- ToDoの論理削除（ソフトデリート）
- 削除されたToDoの復元
- ToDoの物理削除（ハードデリート）

## プロジェクト構成

```
todo-api/
├── main.go           # アプリケーションのエントリーポイント
├── models/           # データモデルとデータベース操作
│   └── todo.go
├── handlers/         # HTTPリクエストハンドラー
│   └── todoHandlers.go
├── go.mod            # モジュール依存関係
└── README.md         # プロジェクト説明（このファイル）
```

## インストール方法

1. リポジトリのクローン:

```bash
git clone https://github.com/yourusername/todo-api.git
cd todo-api
```

2. 依存関係のインストール:

```bash
go mod download
```

3. アプリケーションの実行:

```bash
go run main.go
```

サーバーは `http://localhost:8080` で起動します。

## API エンドポイント

| メソッド | URL                  | 説明                                |
|----------|----------------------|-----------------------------------|
| GET      | /todos               | 全てのToDo項目を取得（論理削除されていないもののみ） |
| GET      | /todos/{id}          | 指定したIDのToDo項目を取得               |
| POST     | /todos               | 新しいToDo項目を作成                   |
| PUT      | /todos/{id}          | 指定したIDのToDo項目を更新               |
| DELETE   | /todos/{id}          | 指定したIDのToDo項目を論理削除            |
| POST     | /todos/{id}/restore  | 論理削除されたToDo項目を復元              |
| DELETE   | /todos/{id}/hard     | 指定したIDのToDo項目を物理削除（完全に削除）     |

## リクエスト例

### ToDo一覧の取得

```bash
curl -X GET http://localhost:8080/todos
```

### 新しいToDoの作成

```bash
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"title": "牛乳を買う"}'
```

### ToDoの更新

```bash
curl -X PUT http://localhost:8080/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "牛乳と卵を買う", "completed": true}'
```

### ToDoの論理削除

```bash
curl -X DELETE http://localhost:8080/todos/1
```

### 削除したToDoの復元

```bash
curl -X POST http://localhost:8080/todos/1/restore
```

### ToDoの物理削除

```bash
curl -X DELETE http://localhost:8080/todos/1/hard
```

## 論理削除について

このAPIでは、`DELETE` メソッドによる通常の削除操作は、実際にはデータを削除せず、`DeletedAt` フィールドに削除日時を設定する「論理削除」を行います。これにより：

- 削除されたデータを必要に応じて復元できる
- 削除履歴を保持できる
- 実際のデータベースからは削除せずに「削除された」状態を表現できる

論理削除されたToDo項目は、一覧取得APIには含まれません。完全に削除したい場合は、ハードデリートエンドポイント `/todos/{id}/hard` を使用します。
