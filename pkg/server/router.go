package server

import (
	"bytes"
	"log"
	"net/http"
	"text/template"
)

type Router struct {
	config *Config
	logger *log.Logger
}

func (ro *Router) Handle(w http.ResponseWriter, r *http.Request) {
	ro.logger.Printf("%s  %s", r.Method, r.URL.Path)

	endpoint := ro.config.Endpoints.Match(r.URL.Path)
	if endpoint == nil {
		if err := UndefinedEndpoint(w, http.StatusNotFound); err != nil {
			ro.logger.Println(err)
		}
		return
	}

	method := endpoint.Methods.Match(r.Method)
	if method == nil {
		if err := UndefinedEndpoint(w, http.StatusMethodNotAllowed); err != nil {
			ro.logger.Println(err)
		}
		return
	}

	if len(method.Queries) == 0 {
		if err := method.Handler.Func(w, method.Handler.Body); err != nil {
			ro.logger.Println(err)
		}
		return
	}

	var body bytes.Buffer
	t, err := template.New("body").Parse(method.Handler.Body.(string))
	if err != nil {
		ro.logger.Println(err)
		return
	}
	qmap := map[string]string{}
	for _, v := range method.Queries {
		qmap[v] = r.URL.Query().Get(v)
	}
	if err = t.Execute(&body, qmap); err != nil {
		ro.logger.Println(err)
		return
	}
	if err = method.Handler.Func(w, body.String()); err != nil {
		ro.logger.Println(err)
	}
}
