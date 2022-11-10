package utils

import "net/http"

func WriteStatusCode(body []byte, w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(body)
}

func Write200(body []byte, w http.ResponseWriter) {
	WriteStatusCode(body, w, http.StatusOK)
}

func Write400(err string, w http.ResponseWriter) {
	WriteStatusCode([]byte(err), w, http.StatusBadRequest)
}

func Write500(err string, w http.ResponseWriter) {
	WriteStatusCode([]byte(err), w, http.StatusInternalServerError)
}
