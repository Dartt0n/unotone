package health_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dartt0n/unotone/controllers/health"
)

func TestController_Health(t *testing.T) {
	t.Parallel()

	ctrl := health.NewController()

	t.Run("should return 200 OK", func(t *testing.T) {
		t.Parallel()

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)

		ctrl.Health(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "OK", w.Body.String())
	})

}
