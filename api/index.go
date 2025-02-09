package main

import (
	"net/http"

	"github.com/Sparker0i/Cactro-Backend-09-Feb-25/internal/config"
	"github.com/Sparker0i/Cactro-Backend-09-Feb-25/internal/server"
)

var router http.Handler

func init() {
	// Load configuration and create the router.
	cfg := config.LoadConfig()
	router = server.NewRouter(cfg)
}

// Handler is the exported function that Vercel invokes for every HTTP request.
func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
