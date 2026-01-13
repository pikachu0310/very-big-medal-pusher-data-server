package main

import (
	_ "embed"
	"fmt"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/handler"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/migration"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/pkg/config"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/internal/repository"
	"github.com/pikachu0310/very-big-medal-pusher-data-server/openapi"
)

//go:embed openapi/openapi.yaml
var openapiYAML []byte

const swaggerHTML = `<!doctype html>
<html>
<head>
  <meta charset="UTF-8">
  <title>Swagger UI</title>
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/swagger-ui-dist@5/swagger-ui.css">
</head>
<body>
  <div id="swagger-ui"></div>
  <script src="https://cdn.jsdelivr.net/npm/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
  <script src="https://cdn.jsdelivr.net/npm/swagger-ui-dist@5/swagger-ui-standalone-preset.js"></script>
  <script>
    window.onload = () => {
      SwaggerUIBundle({
        dom_id: '#swagger-ui',
        urls: [
          {url: '%s', name: 'local'},
          {url: 'https://push.trap.games/api/openapi.yaml', name: 'prod'},
          {url: 'https://push-test.trap.games/api/openapi.yaml', name: 'staging'}
        ],
        layout: 'StandaloneLayout',
        presets: [SwaggerUIBundle.presets.apis, SwaggerUIStandalonePreset],
      });
    };
  </script>
</body>
</html>`

func main() {
	e := echo.New()

	swagger, err := openapi.GetSwagger()
	if err != nil {
		e.Logger.Fatal("Error loading swagger spec\n: %s", err)
	}

	baseURL := "/api"
	swagger.Servers = openapi3.Servers{&openapi3.Server{URL: baseURL}}

	// middlewares
	e.Use(middleware.Recover())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogMethod:   true,
		LogURI:      true,
		LogRemoteIP: true,
	}))
	//e.Use(oapimiddleware.OapiRequestValidator(swagger))

	// connect to database
	db, err := sqlx.Connect("mysql", config.MySQL().FormatDSN())
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer func() {
		_ = db.Close()
	}()

	// migrate tables
	if err := migration.MigrateTables(db.DB); err != nil {
		e.Logger.Fatal(err)
	}

	// setup repository
	repo := repository.New(db)

	// setup routes
	h := handler.New(repo)
	openapi.RegisterHandlersWithBaseURL(e, h, baseURL)

	// expose OpenAPI and Swagger UI
	e.GET(baseURL+"/openapi.yaml", func(c echo.Context) error {
		return c.Blob(http.StatusOK, "application/yaml", openapiYAML)
	})
	e.GET("/swagger", func(c echo.Context) error {
		return c.Redirect(http.StatusFound, "/swagger/index.html")
	})
	e.GET("/swagger/index.html", func(c echo.Context) error {
		html := fmt.Sprintf(swaggerHTML, baseURL+"/openapi.yaml")
		return c.HTML(http.StatusOK, html)
	})

	e.Logger.Fatal(e.Start(config.AppAddr()))
}
