package main

import (
	"net/http"
	"task-manager/db"
	"task-manager/handlers"
	"task-manager/models"

	"github.com/go-chi/chi/v5"
)

func main() {
	db.Init()
	db.DB.AutoMigrate(&models.Task{})

	r := chi.NewRouter()

	r.Post("/tasks", handlers.CreateTask)
	r.Get("/tasks", handlers.GetTasks)
	r.Delete("/tasks/{id}", handlers.DeleteTask)

	http.ListenAndServe(":8080", r)
}
