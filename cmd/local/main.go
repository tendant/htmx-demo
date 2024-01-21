package main

import (
	"html/template"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/tendant/chi-demo/server"
	"github.com/tendant/htmx-demo/idm"
)

func main() {
	var cfg server.Config
	cleanenv.ReadEnv(&cfg)
	s := server.Default(cfg)
	server.Routes(s.R)

	tmplFile := "login.tmpl"
	tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}

	handle := idm.Handle{
		T: tmpl,
	}

	handle.Routes(s.R)

	s.Run()

}
