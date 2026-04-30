package handler

import (
	"strings"
	"unicode/utf8"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RequestLogMiddleware(baseURL string) echo.MiddlewareFunc {
	skippedPaths := map[string]struct{}{
		strings.TrimRight(baseURL, "/") + "/ping":          {},
		strings.TrimRight(baseURL, "/") + "/v4/statistics": {},
	}

	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:    true,
		LogRemoteIP:  true,
		LogUserAgent: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if shouldSkipRequestLog(c, skippedPaths) {
				return nil
			}

			api := formatAPIPath(baseURL, c)
			userName := requestLogUserName(c)
			userAgent := v.UserAgent
			if userAgent == "" {
				userAgent = "-"
			}
			ip := v.RemoteIP
			if ip == "" {
				ip = "-"
			}

			c.Logger().Infof("[%s] status=%d user_name=%q user_agent=%q ip=%s", api, v.Status, userName, userAgent, ip)
			return nil
		},
	})
}

func shouldSkipRequestLog(c echo.Context, skippedPaths map[string]struct{}) bool {
	if _, ok := skippedPaths[c.Request().URL.Path]; ok {
		return true
	}
	if _, ok := skippedPaths[c.Path()]; ok {
		return true
	}
	return false
}

func formatAPIPath(baseURL string, c echo.Context) string {
	path := c.Path()
	if path == "" {
		path = c.Request().URL.Path
	}
	if baseURL != "" && strings.HasPrefix(path, baseURL) {
		path = strings.TrimPrefix(path, baseURL)
	}
	if path == "" {
		path = "/"
	}
	path = strings.ReplaceAll(path, ":user_id", "{user_id}")
	return path
}

func requestLogUserName(c echo.Context) string {
	raw := c.Param("user_id")
	if raw == "" {
		raw = c.QueryParam("user_id")
	}
	if raw == "" {
		return "-"
	}
	decoded, err := decodeUserIDParam(raw)
	if err != nil || decoded == "" {
		return raw
	}
	if !utf8.ValidString(decoded) {
		return raw
	}
	return decoded
}
