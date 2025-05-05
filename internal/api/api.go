package api

import (
	"encoding/json"
	"net/http"
	"step1_simple_api/internal/types"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
)

var Taks float64

type API struct {
	router *chi.Mux
}

func New() *API {
	api := &API{router: chi.NewRouter()}
	api.registerEndpoints()

	return api
}

func (api *API) Serve() error {
	srv := &http.Server{
		Addr:    			":8080",
		Handler: 			api.router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	return srv.ListenAndServe()
}

func (api *API) registerEndpoints() {
	api.router.Get("/api/v1/calculates", api.lastResult)
	api.router.Post("/api/v1/calculates", api.operation)
}

func (api *API) lastResult(w http.ResponseWriter, r *http.Request) {
	api.WriteJSON(w, r, Taks)
}

func (api *API) operation(w http.ResponseWriter, r *http.Request) {
	var req types.Body
	_ = json.NewDecoder(r.Body).Decode(&req)

	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()

	Taks = req.Task
}
