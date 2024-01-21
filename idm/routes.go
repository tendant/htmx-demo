package idm

import (
	"html/template"
	"log/slog"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Handle struct {
	T *template.Template
}

func (h *Handle) Login(w http.ResponseWriter, r *http.Request) {

	s := new(strings.Builder)
	h.T.ExecuteTemplate(s, "login.tmpl", "")
	slog.Info("login:", "s", s.String())
	render.HTML(w, r, s.String())
}

func (h *Handle) Routes(r *chi.Mux) {
	r.Get("/local/login", h.Login)
}
