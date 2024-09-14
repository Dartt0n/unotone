package health

import (
	"log/slog"
	"net/http"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Health(w http.ResponseWriter, r *http.Request) {
	slog.Info("serving health check")
	w.Write([]byte("OK"))
}
