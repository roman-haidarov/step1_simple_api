package api

import (
	"encoding/json"
	"net/http"
	"step1_simple_api/internal/tasks"
	"step1_simple_api/internal/types"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type API struct {
	router *chi.Mux
	tasks  *tasks.Service
}

func New(tasks *tasks.Service) *API {
	api := &API{
		router: chi.NewRouter(),
		tasks:  tasks,
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
	tasks, err := api.tasks.ObjectsTasks()

	if err != nil {
		api.WriteError(w, r, "Internal server error", http.StatusInternalServerError)
		return
	}

	api.WriteJSON(w, r, tasks, http.StatusOK)
}

func (api *API) getTask(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")
	task, err := api.tasks.ObjectTask(uuid)

	if err != nil {
		api.WriteError(w, r, "Task not found", http.StatusNotFound)
		return
	}

	api.WriteJSON(w, r, task, http.StatusOK)
}

func (api *API) createTask(w http.ResponseWriter, r *http.Request) {
	var req types.Task
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		api.WriteError(w, r, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	req.UUID = uuid.NewString()

	api.tasks.CreateTask(req)
	api.WriteJSON(w, r, req, http.StatusCreated)
}

func (api *API) updateTask(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")
	updatedTask := types.Task{UUID: uuid}

	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		api.WriteError(w, r, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	_, err := api.tasks.ObjectTask(uuid)

	if err != nil {
		api.WriteError(w, r, "Task not found", http.StatusBadRequest)
		return
	}

	if err = api.tasks.UpdateTask(updatedTask); err != nil {
		api.WriteError(w, r, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	api.WriteJSON(w, r, updatedTask, http.StatusOK)
}

func (api *API) destroyTask(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")

	if err := api.tasks.DestroyTask(uuid); err != nil {
		api.WriteError(w, r, "Task not found", http.StatusNotFound)
		return
	}

	api.WriteJSON(w, r, map[string]string{uuid: uuid}, http.StatusNoContent)
}
