package main

import (
	"net/http"
	"github.com/tomcoakes/movie-app/handlers"
)

func main() {
	http.HandleFunc("/healthcheck", handlers.HealthcheckHandler)
	http.ListenAndServe(":8080", nil)
}
