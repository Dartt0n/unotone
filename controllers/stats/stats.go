package stats

import (
	"math/rand"
	"net/http"
	"strconv"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) CountImages(w http.ResponseWriter, r *http.Request) {
	numberOfImages := rand.Intn(150) + 20

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(numberOfImages)))
}
