package main

import (
	"net/http"

	"github.com/AbilashKriz/bookings/pkg/config"
	"github.com/AbilashKriz/bookings/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Route(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/where", handlers.Repo.Where)
	mux.Get("/about", handlers.Repo.About)

	return mux

}
