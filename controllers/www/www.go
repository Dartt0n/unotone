package www

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/dartt0n/unotone/controllers/www/components"
)

type Controller struct {
	templ  http.Handler
	static http.Handler
}

func NewController(staticDir string) *Controller {
	return &Controller{
		templ:  templ.Handler(components.Page()),
		static: http.FileServer(http.Dir(staticDir)),
	}
}

func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	slog.Info("serving web page", "path", r.URL.EscapedPath())
	c.templ.ServeHTTP(w, r)
}

func (c *Controller) Static(w http.ResponseWriter, r *http.Request) {
	slog.Info("serving static file", "path", r.URL.EscapedPath())
	c.static.ServeHTTP(w, r)
}
