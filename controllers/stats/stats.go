package stats

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/dartt0n/unotone/controllers/web/components"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) GetImageCountPretty(w http.ResponseWriter, r *http.Request) {
	n := rand.Intn(350) + 20

	components.RaisingNumbers(strconv.Itoa(n)).Render(r.Context(), w)
}

func (c *Controller) GetImageCountRaw(w http.ResponseWriter, r *http.Request) {
	n := rand.Intn(350) + 20

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(n)))
}
