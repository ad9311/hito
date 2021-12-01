package main

import (
	"net/http"
)

func serve() *http.Server {
	serve := &http.Server{
		Addr:    config.PortNumber,
		Handler: routes(),
	}
	return serve
}
