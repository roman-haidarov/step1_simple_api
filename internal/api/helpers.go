package api

import (
	"encoding/json"
	"net/http"
)

func (api *API) WriteJSON(w http.ResponseWriter, r *http.Request, response any) {
	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(response)
}
