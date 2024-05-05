package handlers

import "net/http"

func SetConfig(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Pong!"))
}
