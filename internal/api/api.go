package api

import (
	"encoding/json"
	"net/http"
	"step1_simple_api/internal/types"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type API struct {
	router *chi.Mux
	mu     sync.Mutex
	tasks  map[string]types.Task
}

func New() *API {
	api := &API{
		router: chi.NewRouter(),
		tasks:  make(map[string]types.Task),
	}
	api.registerEndpoints()

	return api
}

func (api *API) Serve() error {
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      api.router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	return srv.ListenAndServe()
}

func (api *API) registerEndpoints() {
	api.router.Get("/api/v1/tasks", api.getTasks)
	api.router.Get("/api/v1/tasks/{uuid}", api.getTask)
	api.router.Post("/api/v1/tasks", api.createTask)
	api.router.Patch("/api/v1/tasks/{uuid}", api.updateTask)
	api.router.Delete("/api/v1/tasks/{uuid}", api.destroyTask)
}

func (api *API) getTasks(w http.ResponseWriter, r *http.Request) {
	api.mu.Lock()
	tasks := make([]types.Task, 0, len(api.tasks))
	for _, task := range api.tasks {
		tasks = append(tasks, task)
	}
	api.mu.Unlock()

	api.WriteJSON(w, r, tasks, http.StatusOK)
}

func (api *API) getTask(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")

	api.mu.Lock()
	task, exists := api.tasks[uuid]
	api.mu.Unlock()

	if exists {
		api.WriteJSON(w, r, task, http.StatusOK)
	} else {
		api.WriteError(w, r, "Task not found", http.StatusNotFound)
	}
}

func (api *API) createTask(w http.ResponseWriter, r *http.Request) {
	var req types.Task
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		api.WriteError(w, r, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	api.mu.Lock()
	req.UUID = uuid.NewString()
	api.tasks[req.UUID] = req
	api.mu.Unlock()

	api.WriteJSON(w, r, req, http.StatusCreated)
}

func (api *API) updateTask(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")
	var updatedTask types.Task

	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		api.WriteError(w, r, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	api.mu.Lock()
	_, exists := api.tasks[uuid]

	if exists {
		updatedTask.UUID = uuid
		api.tasks[uuid] = updatedTask
	}
	api.mu.Unlock()

	if exists {
		api.WriteJSON(w, r, updatedTask, http.StatusOK)
	} else {
		api.WriteError(w, r, "Task not found", http.StatusNotFound)
	}
}

func (api *API) destroyTask(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")

	api.mu.Lock()
	_, exists := api.tasks[uuid]
	if exists {
		delete(api.tasks, uuid)
	}
	api.mu.Unlock()

	if exists {
		api.WriteNoContent(w)
	} else {
		api.WriteError(w, r, "Task not found", http.StatusNotFound)
	}
}
