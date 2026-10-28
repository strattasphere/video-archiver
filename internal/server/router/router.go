package router

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/strattasphere/video-archiver/components"
	"github.com/strattasphere/video-archiver/internal/dist"
	"github.com/strattasphere/video-archiver/internal/server/handler"
)

func New() http.Handler {
	h := handler.New()

	mux := http.NewServeMux()
	addRoutes(mux, h)

	return http.Handler(mux)
}

func addRoutes(m *http.ServeMux, h *handler.Handler) {
	// maintenance routes
	m.HandleFunc(newPath(http.MethodGet, "/health"), h.Health)

	// templ routes
	m.Handle(newPath(http.MethodGet, "/"), templ.Handler(components.Home()))

	// api routes
	m.Handle(newPath(http.MethodGet, "/assets/"), http.FileServer(http.FS(dist.AssetsDir)))
}

func newPath(method, path string) string {
	return method + " " + path
}
