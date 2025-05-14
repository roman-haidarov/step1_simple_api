package api

import (
	"encoding/json"
	"net/http"
)

func (api *API) WriteJSON(w http.ResponseWriter, r *http.Request, response any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(response)
}

func (api *API) WriteError(w http.ResponseWriter, r *http.Request, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	error := map[string]string{"error": message}
	_ = json.NewEncoder(w).Encode(error)
}

func (api *API) WriteNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
