package truth_service

import (
	"fmt"
	"log"
	"net/http"
)

func Run() {
	httpRouter()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// simple router
func httpRouter() {
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/alive", aliveHandler)
}

// handlers for routes
func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func aliveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "youbetcha")
}
