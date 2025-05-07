# Task Manager на Go

Простой REST API для управления задачами, написанный на Go с использованием GORM и PostgreSQL.

## Функции

- Создание задач (POST /tasks)
- Получение списка задач (GET /tasks)
- Удаление задачи по ID (DELETE /tasks/{id})

## Стек

- Go 1.21+
- PostgreSQL
- GORM
- Chi router


## Установка и запуск

1. Установите Go и PostgreSQL
2. Клонируйте репозиторий и перейдите в папку:

```bash
git clone git@github.com:yourname/task-manager.git
cd task-manager
```

3. Установите зависимости:

```bash
go mod tidy
```

4. Создайте базу данных `task_manager` в PostgreSQL и проверьте строку подключения в `db/db.go`:

```go
dsn := "host=localhost user=postgres password=postgres dbname=task_manager port=5432 sslmode=disable"
```

5. Запустите сервер:

```bash
go run main.go
```

---

## Примеры тестирования API

Можно тестировать с помощью `curl`:

### Создать задачу

```bash
curl -X POST http://localhost:8080/tasks \
-H "Content-Type: application/json" \
-d '{
  "title": "Доделать проект",
  "description": "Go + PostgreSQL",
  "is_done": false,
  "user_id": 1
}'
```

### Получить список задач

```bash
curl http://localhost:8080/tasks
```

### Удалить задачу

```bash
curl -X DELETE http://localhost:8080/tasks/1
```

---

## Пример JSON-ответа

```json
{
  "ID": 1,
  "Title": "Доделать проект",
  "Description": "Go + PostgreSQL",
  "IsDone": false,
  "DueDate": "0001-01-01T00:00:00Z",
  "UserID": 1
}
```

---

Проект создан в учебных целях.
