package BookHttpMiddleware_test

import (
	"github.com/labstack/echo/v4"
	"kom-product-service/adapters/product/delivery/http/middleware"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCORS(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest(echo.GET, "/", nil)
	response := httptest.NewRecorder()
	context := e.NewContext(request, response)
	middleware := BookHttpMiddleware.New()

	handler := middleware.CORS(echo.HandlerFunc(func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	}))

	err := handler(context)

	assert.Equal(t, "*", response.Header().Get("Access-Control-Allow-Origin"))
	require.NoError(t, err)
}
