package negronihealth

import (
	"net/http"
)

type Option func(h *Health)

func Path(path string) Option {
	return func(h *Health) {
		h.path = path
	}
}

func Status(status int) Option {
	return func(h *Health) {
		h.status = status
	}
}

type Health struct {
	path   string
	status int
}

func New(options ...Option) *Health {
	var h Health

	for _, option := range options {
		option(&h)
	}

	return &h
}

func (h *Health) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if r.URL.Path == h.path {
		rw.WriteHeader(h.status)
		return
	}

	next(rw, r)
}
