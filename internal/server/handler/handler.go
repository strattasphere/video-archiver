package handler

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/strattasphere/video-archiver/components"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Hello2() *templ.ComponentHandler {
	return templ.Handler(components.Hello("Rufus"))
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) *templ.ComponentHandler {

	return templ.Handler(components.Home())
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("service alive"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("failed to write body: %v", err)
		return
	}
}
