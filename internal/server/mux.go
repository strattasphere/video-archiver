package server

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/strattasphere/video-archiver/components"
)

type serveMux struct {
	*http.ServeMux
}

func NewServeMux() *serveMux {
	return &serveMux{http.NewServeMux()}
}

func (m *serveMux) AddRoutes() {
	m.Handle("GET /health", m.healthHandler())
	m.Handle("GET /hello", templ.Handler(components.hello("Luke")))
}
