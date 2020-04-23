package controllers

import (
	"github.com/Gusarov2k/list-of-expenses/internal/postgres"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestGetSpent(t *testing.T) {
	c := postgres.NewClient()

	if err := c.Open(postgres.PostgresSys); err != nil {
		panic("failed to connect database")
	}
	c.Schema()
	defer func() { c.Close() }()

	e := Router{}
	e.Handler(c.Db)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/spent/2", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	assert.Equal(t, http.StatusOK, rec.Code)
}
