package api

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"step1_simple_api/internal/tasks"
	"step1_simple_api/internal/types"
	generatedTasks "step1_simple_api/internal/web/tasks"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type API struct {
	router *chi.Mux
	tasks  *tasks.Service
}

func (api *API) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := api.tasks.GetTasks()
	if err != nil {
		api.WriteError(w, r, "Internal server error", http.StatusInternalServerError)
		return
	}

	response := make([]generatedTasks.Task, len(tasks))
	for i, task := range tasks {
		response[i] = api.convertToGeneratedTask(task)
	}

	api.WriteJSON(w, r, response, http.StatusOK)
}

func (api *API) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req generatedTasks.CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		api.WriteError(w, r, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	task := types.Task{
		UUID:        uuid.NewString(),
		Description: req.Description,
		IsDone:      req.IsDone != nil && *req.IsDone,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if _, err := api.tasks.CreateTask(task); err != nil {
		api.WriteError(w, r, "Failed to create task", http.StatusBadRequest)
		return
	}

	response := api.convertToGeneratedTask(task)
	api.WriteJSON(w, r, response, http.StatusCreated)
}

func (api *API) GetTask(w http.ResponseWriter, r *http.Request, uuid openapi_types.UUID) {
	task, err := api.tasks.GetTask(uuid.String())
	if err != nil {
		api.WriteError(w, r, "Task not found", http.StatusNotFound)
		return
	}

	response := api.convertToGeneratedTask(task)
	api.WriteJSON(w, r, response, http.StatusOK)
}

func (api *API) UpdateTask(w http.ResponseWriter, r *http.Request, uuid openapi_types.UUID) {
	var req generatedTasks.UpdateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		api.WriteError(w, r, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	existingTask, err := api.tasks.GetTask(uuid.String())
	if err != nil {
		api.WriteError(w, r, "Task not found", http.StatusNotFound)
		return
	}

	updatedTask := existingTask
	if req.Description != nil {
		updatedTask.Description = *req.Description
	}
	if req.IsDone != nil {
		updatedTask.IsDone = *req.IsDone
	}
	updatedTask.UpdatedAt = time.Now()

	if err = api.tasks.UpdateTask(updatedTask); err != nil {
		api.WriteError(w, r, "Failed to update task", http.StatusInternalServerError)
		return
	}

	response := api.convertToGeneratedTask(updatedTask)
	api.WriteJSON(w, r, response, http.StatusOK)
}

func (api *API) DestroyTask(w http.ResponseWriter, r *http.Request, uuid openapi_types.UUID) {
	if err := api.tasks.DestroyTask(uuid.String()); err != nil {
		api.WriteError(w, r, "Task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (api *API) convertToGeneratedTask(task types.Task) generatedTasks.Task {
	parsedUUID, err := uuid.Parse(task.UUID)
	if err != nil {
		parsedUUID = uuid.New()
	}

	var uuidObj openapi_types.UUID
	copy(uuidObj[:], parsedUUID[:])
	
	generatedTask := generatedTasks.Task{
		Uuid:        uuidObj,
		Description: task.Description,
		IsDone:      task.IsDone,
	}

	if !task.CreatedAt.IsZero() {
		generatedTask.CreatedAt = &task.CreatedAt
	}
	if !task.UpdatedAt.IsZero() {
		generatedTask.UpdatedAt = &task.UpdatedAt
	}
	if task.DeletedAt != nil {
		generatedTask.DeletedAt = task.DeletedAt
	}

	return generatedTask
}

func New(tasks *tasks.Service) *API {
	api := &API{
		router: chi.NewRouter(),
		tasks:  tasks,
	}

	handler := generatedTasks.HandlerFromMuxWithBaseURL(api, api.router, "/api/v1")
	api.router.Mount("/", handler)

	return api
}

func (api *API) Serve(ctx context.Context) error {
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      api.router,
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	return srv.ListenAndServe()
}

func (api *API) WriteJSON(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {}
}

func (api *API) WriteError(w http.ResponseWriter, r *http.Request, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	errorResponse := generatedTasks.Error{
		Message: message,
	}
	json.NewEncoder(w).Encode(errorResponse)
}
