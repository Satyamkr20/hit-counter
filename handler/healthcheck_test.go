package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHandler_HealthCheck(t *testing.T) {
	assert := assert.New(t)
	e := echo.New()

	s, err := miniredis.Run()
	assert.NoError(err)
	defer s.Close()

	h, err := NewHandler([]string{s.Addr()})
	assert.NoError(err)

	r := httptest.NewRequest("GET", "http://localhost:8080", nil)
	w := httptest.NewRecorder()

	ectx := e.NewContext(r, w)
	err = h.HealthCheck(ectx)
	assert.NoError(err)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(http.StatusOK, resp.StatusCode)
	assert.Equal("health check!", string(body))
}
