package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

//adds csrf protection to all request

func NoSurf(handler http.Handler) http.Handler {
	csrfHandler := nosurf.New(handler)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		Secure:   app.InProduction,
	})
	return csrfHandler
}

//Loads and saves the session on every request

func SessionLoad(handler http.Handler) http.Handler {
	return session.LoadAndSave(handler)
}
