package main

import (
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/dartt0n/unotone/controllers/health"
	"github.com/dartt0n/unotone/controllers/stats"
	"github.com/dartt0n/unotone/controllers/web"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	Addr      string `koanf:"addr" validate:"required" json:"addr"`
	StaticDir string `koanf:"static_dir" validate:"required" json:"static_dir"`
	Debug     bool   `koanf:"debug" validate:"required" json:"debug"`
}

func main() {
	// read config
	conf := Config{
		// default options
		Addr:      ":8080",
		StaticDir: "/var/www/static",
		Debug:     false,
	}

	k := koanf.New(".")
	k.Load(env.Provider("UNOTONE_", ".", func(s string) string {
		s = strings.TrimPrefix(s, "UNOTONE_")
		s = strings.ToLower(s)
		return s
	}), nil)
	k.Unmarshal("", &conf)

	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(&conf); err != nil {
		slog.Info("config validation failed", "err", err)
		return
	}

	// setup logger
	if conf.Debug {
		slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		})))
	} else {
		slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelInfo,
			AddSource: true,
		})))
	}

	// setup router
	r := chi.NewMux()

	healthCtrl := health.NewController()
	webCtrl := web.NewController(conf.StaticDir)
	statsCtrl := stats.NewController()

	r.Group(func(r chi.Router) {
		r.Get("/health", healthCtrl.Health)

		r.Get("/", webCtrl.Index)
		r.Get("/*", webCtrl.Static)

		r.Get("/api/stats/count/images", statsCtrl.CountImages)
	})

	// run
	slog.Info("starting server", "config", conf)
	http.ListenAndServe(conf.Addr, r)
}
