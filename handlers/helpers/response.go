package helpers

import "net/http"

func SendError(w http.ResponseWriter, response []byte) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(response)
}

func Send(w http.ResponseWriter, response []byte) {
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
