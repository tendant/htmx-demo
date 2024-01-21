package idm

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Handle struct {
	T *template.Template
}

func (h *Handle) Login(w http.ResponseWriter, r *http.Request) {

	h.T.ExecuteTemplate(w, "login", nil)
	render.HTML(w, r, "")
}

func (h *Handle) Routes(r *chi.Mux) {
	r.Get("/local/login", h.Login)
}
