package alive

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
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/looker", lookerHandler)
}

// handlers for routes
func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func aliveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "youbetcha")
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{ alive: true }")
}

func lookerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hey there good looker")
}
