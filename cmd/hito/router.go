package main

import (
	"net/http"

	"github.com/ad9311/hito/internal/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const (
	home   = "/"
	login  = "/login"
	logout = "/logout"

	apiLandmarks = "/api/v1/landmarks"
	users        = "/api/v1/current-user"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(sessionsLoad)
	mux.Use(newCsrf)

	mux.Get(home, handler.Home)
	mux.Post(home, handler.PostHome)

	mux.Get(login, handler.Login)
	mux.Post(login, handler.PostLogin)
	mux.Get(logout, handler.Logout)

	mux.Get(apiLandmarks, handler.Landmarks)
	mux.Post(users, handler.PostCurrentUser)

	fileServer := http.FileServer(http.Dir("./web/static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux
}
