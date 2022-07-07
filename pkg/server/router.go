package server

import (
	"log"
	"net/http"
)

type Router struct {
	config *Config
	logger *log.Logger
}

func (ro *Router) Handle(w http.ResponseWriter, r *http.Request) {
	ro.logger.Printf("%s  %s", r.Method, r.URL.Path)

	endpoint := ro.config.Endpoints.Match(r.URL.Path)
	if endpoint == nil {
		UndefinedEndpoint(w, http.StatusNotFound)
		return
	}

	handler := endpoint.Methods.Match(r.Method)
	if handler == nil {
		UndefinedEndpoint(w, http.StatusMethodNotAllowed)
		return
	}

	handler.Func(w, r)
}
