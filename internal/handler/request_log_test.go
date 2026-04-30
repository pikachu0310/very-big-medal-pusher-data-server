package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	echolog "github.com/labstack/gommon/log"
)

func TestRequestLogMiddlewareFormatsAndRedacts(t *testing.T) {
	var logs bytes.Buffer
	e := echo.New()
	e.Logger.SetOutput(&logs)
	e.Logger.SetLevel(echolog.INFO)
	e.Logger.SetHeader("")
	e.Use(RequestLogMiddleware("/api"))
	e.GET("/api/v4/users/:user_id/data", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	req := httptest.NewRequest(http.MethodGet, "/api/v4/users/pikachu0310/data?data=payload&sig=secret", nil)
	req.Header.Set("User-Agent", "unity")
	req.Header.Set(echo.HeaderXForwardedFor, "192.1.1.1")
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	got := logs.String()
	want := `[/v4/users/{user_id}/data] status=200 user_name="pikachu0310" user_agent="unity" ip=192.1.1.1`
	if !strings.Contains(got, want) {
		t.Fatalf("log output missing expected entry:\nwant contains: %s\ngot: %s", want, got)
	}
	if strings.Contains(got, "payload") || strings.Contains(got, "secret") {
		t.Fatalf("log output leaked query data or signature: %s", got)
	}
}

func TestRequestLogMiddlewareSkipsV4Statistics(t *testing.T) {
	var logs bytes.Buffer
	e := echo.New()
	e.Logger.SetOutput(&logs)
	e.Logger.SetLevel(echolog.INFO)
	e.Logger.SetHeader("")
	e.Use(RequestLogMiddleware("/api"))
	e.GET("/api/v4/statistics", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	req := httptest.NewRequest(http.MethodGet, "/api/v4/statistics", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if logs.String() != "" {
		t.Fatalf("expected no log output for /v4/statistics, got: %s", logs.String())
	}
}

func TestRequestLogMiddlewareSkipsPing(t *testing.T) {
	var logs bytes.Buffer
	e := echo.New()
	e.Logger.SetOutput(&logs)
	e.Logger.SetLevel(echolog.INFO)
	e.Logger.SetHeader("")
	e.Use(RequestLogMiddleware("/api"))
	e.GET("/api/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	req := httptest.NewRequest(http.MethodGet, "/api/ping", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if logs.String() != "" {
		t.Fatalf("expected no log output for /ping, got: %s", logs.String())
	}
}
