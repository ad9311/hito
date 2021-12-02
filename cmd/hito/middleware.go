package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

func newCsrf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   config.Production,
		SameSite: http.SameSiteLaxMode,
	})
	csrfHandler.ExemptPath(users)
	return csrfHandler
}

func sessionsLoad(next http.Handler) http.Handler {
	return config.Session.LoadAndSave(next)
}
