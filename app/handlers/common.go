package handlers

import (
	"encoding/json"
	"net/http"
)

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	respondJson(w, http.StatusOK, "Healthy!")
}

func respondJson(w http.ResponseWriter, status int, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		handleServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func handleServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
