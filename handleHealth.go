package main

import (
	"net/http"
)

func HandleReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, struct {
		Status string `json:"status"`
}{
		Status: "alive",
})
}
