package main

import (
	"net/http"

	"github.com/dartt0n/unotone/controllers/health"
	"github.com/dartt0n/unotone/controllers/www"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewMux()

	healthCtrl := health.NewController()
	wwwCtrl := www.NewController("./controllers/www")

	r.Group(func(r chi.Router) {
		r.Get("/health", healthCtrl.Health)

		r.Get("/static/*", wwwCtrl.Static)
		r.Get("/", wwwCtrl.Index)
	})

	http.ListenAndServe(":8080", r)
}
