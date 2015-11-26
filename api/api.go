package api

import (
	"net/http"

	"github.com/zenazn/goji/web"
	gojimw "github.com/zenazn/goji/web/middleware"
)

// New returns a new Goji Mux handler to process HTTP requests
func New() http.Handler {
	r := web.New()

	// Mount middleware
	r.Use(gojimw.EnvInit)

	// Mount request handlers
	r.Get("/", defaultHandler)
	r.Get("/version", defaultHandler)

	return r
}
