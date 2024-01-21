package idm

import (
	"html/template"
	"log/slog"
	"net/http"
	"strings"

	"github.com/ggicci/httpin"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Handle struct {
	T *template.Template
}

func (h *Handle) Routes(r *chi.Mux) {
	r.Get("/local/login", h.Login)
	r.With(httpin.NewInput(LoginInput{})).Post("/local/login", h.LoginPost)
}

func (h *Handle) Login(w http.ResponseWriter, r *http.Request) {

	s := new(strings.Builder)
	h.T.ExecuteTemplate(s, "login.tmpl", "")
	slog.Info("login:", "s", s.String())
	render.HTML(w, r, s.String())
}

type LoginInput struct {
	Email    string `in:"form=email"`
	Password string `in:"form=password"`
}

func (h *Handle) LoginPost(w http.ResponseWriter, r *http.Request) {
	// query := r.Context().Value(httpin.Input).(*LoginInput)
	w.Header().Add("HX-Redirect", "http://localhost:4000/")
	// w.Header()["HX-Redirect"] = []string{"https://google.com/"}
	// w.Header()["HX-Refresh"] = []string{"true"}
	// w.Header()["HX-Push-Url"] = []string{"http://localhost:4000/"}
	// http.Redirect(w, r, "/", http.StatusFound)
	http.Redirect(w, r, "http://localhost:4000/", http.StatusOK)
}
