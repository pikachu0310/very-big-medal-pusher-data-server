package main

import (
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

func main() {
	e := echo.New()

	swagger, err := openapi.GetSwagger()
	if err != nil {
		e.Logger.Fatal("Error loading swagger spec\n: %s", err)
	}

	baseURL := ""
	swagger.Servers = openapi3.Servers{&openapi3.Server{URL: baseURL}}

	// middlewares
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	//e.Use(oapimiddleware.OapiRequestValidator(swagger))

	// connect to database
	db, err := sqlx.Connect("mysql", config.MySQL().FormatDSN())
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer db.Close()

	// migrate tables
	if err := migration.MigrateTables(db.DB); err != nil {
		e.Logger.Fatal(err)
	}

	// setup repository
	repo := repository.New(db)

	// setup routes
	h := handler.New(repo)
	openapi.RegisterHandlersWithBaseURL(e, h, baseURL)
	handler.GlobalSecret = config.GetSecretKey()

	e.Logger.Fatal(e.Start(config.AppAddr()))
}
