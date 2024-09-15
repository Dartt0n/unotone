package main

import (
	"net/http"

	"github.com/dartt0n/unotone/controllers/health"
	"github.com/dartt0n/unotone/controllers/stats"
	"github.com/dartt0n/unotone/controllers/web"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewMux()

	healthCtrl := health.NewController()
	webCtrl := web.NewController("./controllers/web/static")
	statsCtrl := stats.NewController()

	r.Group(func(r chi.Router) {
		r.Get("/health", healthCtrl.Health)

		r.Get("/", webCtrl.Index)
		r.Get("/*", webCtrl.Static)

		r.Get("/api/stats/count/images", statsCtrl.CountImages)
	})

	http.ListenAndServe(":8080", r)
}
