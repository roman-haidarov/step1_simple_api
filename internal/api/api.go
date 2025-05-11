package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"step1_simple_api/internal/types"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
)

type API struct {
	router *chi.Mux
}

var Task string

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
	api.router.Get("/api/v1/tasks", api.getTask)
	api.router.Post("/api/v1/tasks", api.createTask)
}

func (api *API) getTask(w http.ResponseWriter, r *http.Request) {
	mess := map[string]string{}
	mess["message"] = fmt.Sprintf("Hello, %v", Task)

	api.WriteJSON(w, r, mess)
}

func (api *API) createTask(w http.ResponseWriter, r *http.Request) {
	var req types.Body
	_ = json.NewDecoder(r.Body).Decode(&req)

	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()

	Task = fmt.Sprintf("%v", req.Task)
	w.WriteHeader(http.StatusCreated)
}
