package main

import (
	"embed"
	"html/template"
	"mime"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/tendant/chi-demo/server"
	"github.com/tendant/htmx-demo/idm"
	"golang.org/x/exp/slog"
)

//go:embed static/*
var distFiles embed.FS

func StaticHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mime.AddExtensionType(".css", "text/css; charset=utf-8")
		f, err := distFiles.Open(strings.TrimPrefix(path.Clean(r.URL.Path), "/"))
		if err == nil {
			defer f.Close()
		}
		if os.IsNotExist(err) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			http.FileServer(http.FS(distFiles)).ServeHTTP(w, r)
		}
	}
}

func main() {
	var cfg server.Config
	cleanenv.ReadEnv(&cfg)
	s := server.Default(cfg)
	server.Routes(s.R)

	tmplFile := "templates/login.tmpl"
	tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	for i, t := range tmpl.Templates() {
		slog.Info("template:", "i", i, "t", t.Name())
	}

	handle := idm.Handle{
		T: tmpl,
	}

	handle.Routes(s.R)

	s.R.Get("/static/css/output.css", StaticHandler())
	s.R.Get("/static/js/htmx.min.js", StaticHandler())
	s.Run()

}
