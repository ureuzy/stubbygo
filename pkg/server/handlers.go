package server

import (
	"fmt"
	"net/http"
)

type Endpoint struct {
	Path    string  `yaml:"path"`
	Methods Methods `yaml:"methods"`
}

type Endpoints []*Endpoint

func (e *Endpoints) Match(param string) *Endpoint {
	for _, res := range *e {
		if param == res.Path {
			return res
		}
	}
	return nil
}

type Method struct {
	Type    string    `yaml:"type"`
	Handler *Response `yaml:"response"`
}

type Methods []*Method

func (m *Methods) Match(method string) *Response {
	for _, v := range *m {
		if method == v.Type {
			return v.Handler
		}
	}
	return nil
}

type Response struct {
	StatusCode  int               `yaml:"status_code"`
	Headers     map[string]string `yaml:"headers"`
	ContentType string            `yaml:"content_type"`
	Body        any               `yaml:"body"`
}

type Responses []*Response

func (r *Response) Func(w http.ResponseWriter, _ *http.Request) {
	for key, value := range r.Headers {
		w.Header().Set(key, value)
	}
	w.WriteHeader(r.StatusCode)
	fmt.Fprint(w, r.Body)
}

func UndefinedEndpoint(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	fmt.Fprint(w, http.StatusText(status))
}
