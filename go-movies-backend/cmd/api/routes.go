package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	// Create router multiplexer
	mux := chi.NewRouter()

	// Log backtrace if error occurs and recover page content
	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	mux.Get("/", app.Home)

	mux.Get("/authenticate", app.authenticate)

	mux.Get("/movies", app.AllMovies)

	return mux
}
