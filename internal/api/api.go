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
}

var Tasks = make(map[string]types.Task)

func New() *API {
	api := &API{router: chi.NewRouter()}
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
	// api.router.Patch("/api/v1/tasks/{uuid}", api.updateTask)
}

func (api *API) getTasks(w http.ResponseWriter, r *http.Request) {
	tasks := make([]types.Task, 0, len(Tasks))
	for _, task := range Tasks {
		tasks = append(tasks, task)
	}

	api.WriteJSON(w, r, tasks)
}

func (api *API) getTask(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")
	task := Tasks[uuid]

	api.WriteJSON(w, r, task)
}

func (api *API) createTask(w http.ResponseWriter, r *http.Request) {
	var req types.Task
	_ = json.NewDecoder(r.Body).Decode(&req)

	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()

	req.UUID = uuid.NewString()
	Tasks[req.UUID] = req
	api.WriteJSON(w, r, req)
}

// func (api *API) updateTask(w http.ResponseWriter, r *http.Request) {
// 	uuid := chi.URLParam(r, "uuid")
// }
