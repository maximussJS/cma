package responses

import "net/http"

func SuccessNull(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("null"))
}

func SuccessJson(w http.ResponseWriter, content string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(content))
}

func SuccessEmpty(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}
