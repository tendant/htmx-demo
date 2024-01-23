package main

import (
	"context"
	"embed"
	"fmt"
	"html/template"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jackc/pgx/v5/pgxpool"
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

type DbConfig struct {
	Host     string
	Port     uint16
	Database string
	User     string
	Password string
}

func (c DbConfig) toDatabaseUrl() string {
	u := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(c.User, c.Password),
		Host:   fmt.Sprintf("%s:%d", c.Host, c.Port),
		Path:   c.Database,
	}
	return u.String()
}

type DbConf interface {
	toDbConfig() DbConfig
}

type DemoDbConfig struct {
	Host     string `env:"HTMX_DEMO_PG_HOST" env-default:"localhost"`
	Port     uint16 `env:"HTMX_DEMO_PG_PORT" env-default:"5432"`
	Database string `env:"HTMX_DEMO_PG_DATABASE" env-default:"htmx_demo_db"`
	User     string `env:"HTMX_DEMO_PG_USER" env-default:"htmx_demo"`
	Password string `env:"HTMX_DEMO_PG_PASSWORD" env-default:"pwd"`
}

func (d DemoDbConfig) toDbConfig() DbConfig {
	return DbConfig{
		Host:     d.Host,
		Port:     d.Port,
		Database: d.Database,
		User:     d.User,
		Password: d.Password,
	}
}

type Config struct {
	ServerConfig server.Config
	DemoDb       DemoDbConfig
}

func main() {
	var cfg Config
	cleanenv.ReadEnv(&cfg)
	s := server.Default(cfg.ServerConfig)
	server.Routes(s.R)

	tmplFile := "templates/login.tmpl"
	tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	for i, t := range tmpl.Templates() {
		slog.Info("template:", "i", i, "t", t.Name())
	}

	var queries *idm.Queries
	pool, err := pgxpool.New(context.Background(), cfg.DemoDb.toDbConfig().toDatabaseUrl())
	if err != nil {
		slog.Error("Failed creating dbpool", "db", cfg.DemoDb.Database, "url", cfg.DemoDb.toDbConfig().toDatabaseUrl())
		os.Exit(-1)
	} else {
		queries = idm.New(pool)
	}

	handle := idm.Handle{
		T: tmpl,
		IdmService: &idm.Service{
			Queries: queries,
		},
	}

	handle.Routes(s.R)

	s.R.Get("/static/css/output.css", StaticHandler())
	s.R.Get("/static/js/htmx.min.js", StaticHandler())
	s.Run()

}
