package idm

import (
	"html/template"
	"log/slog"
	"net/http"
	"net/mail"
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

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (h *Handle) LoginPost(w http.ResponseWriter, r *http.Request) {
	form := r.Context().Value(httpin.Input).(*LoginInput)
	if len(form.Email) == 0 || len(form.Password) == 0 || !valid(form.Email) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	correct, err := VerifyPassword(form.Email, form.Password)
	if err != nil || !correct {
		slog.Warn("Failed verifying password", "email", form.Email, "err", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	} else {
		w.Header().Add("HX-Redirect", "/")
		// http.Redirect(w, r, "http://localhost:4000/", http.StatusOK)
		http.Redirect(w, r, "/", http.StatusOK)

	}
}
