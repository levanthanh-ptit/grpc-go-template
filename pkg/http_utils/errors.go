package http_utils

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Code    uint32        `json:"code,omitempty"`
	Message string        `json:"message,omitempty"`
	Details []interface{} `json:"details,omitempty"`
}

func JSONError(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}
