package api

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"step1_simple_api/internal/tasks"
	"step1_simple_api/internal/users"
	generatedTasks "step1_simple_api/internal/web/tasks"
	generatedUsers "step1_simple_api/internal/web/users"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type API struct {
	router *chi.Mux
	tasks  *tasks.Service
	users  *users.Service
}

func New(tasks *tasks.Service, users *users.Service) *API {
	api := &API{
		router: chi.NewRouter(),
		tasks:  tasks,
		users: 	users,
	}

	generatedTasks.HandlerFromMux(api, api.router)
	generatedUsers.HandlerFromMux(api, api.router)

	api.router.Mount("/api/v1", api.router)

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

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Debug().Err(err).Msg("Failed to encode JSON response")
	}
}

func (api *API) WriteError(w http.ResponseWriter, r *http.Request, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	errorResponse := generatedTasks.Error{
		Message: message,
	}
	if err := json.NewEncoder(w).Encode(errorResponse); err != nil {
		log.Debug().Err(err).Msg("Failed to encode error response")
	}
}
