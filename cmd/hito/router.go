package main

import (
	"net/http"

	"github.com/ad9311/hito/internal/handler"
	"github.com/go-chi/chi"
)

const (
	home   = "/"
	login  = "/login"
	logout = "/logout"

	apiLandmarks = "/api/v1/landmarks"
	users        = "/api/v1/users"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	// mux.Use(middleware.Recoverer)
	mux.Use(sessionsLoad)
	mux.Use(newCsrf)

	mux.Get(home, handler.Home)
	mux.Post(home, handler.PostHome)
	mux.Get(login, handler.Login)
	mux.Post(login, handler.PostLogin)
	mux.Get(logout, handler.Logout)

	mux.Get(apiLandmarks, handler.Landmarks)
	mux.Get(users, handler.Users)

	fileServer := http.FileServer(http.Dir("./web/static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux
}
