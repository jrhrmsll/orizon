package http

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/jrhrmsll/orizon"
)

const StatusCodeContextCanceled = 499

type Application struct {
	ctx    context.Context
	logger *log.Logger
	srv    *http.Server
}

func NewApplication(
	ctx context.Context,
	logger *log.Logger,
	intervalService orizon.IntervalService,
) *Application {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	prometheusMiddleware := echoprometheus.NewMiddleware("http")
	e.GET("/metrics", echoprometheus.NewHandler())

	api := e.Group("api/v1")
	api.Use(prometheusMiddleware)

	var intervalController = NewIntervalController(intervalService)

	api.POST("/intervals", intervalController.Index)

	return &Application{
		ctx:    ctx,
		logger: logger,
		srv: &http.Server{
			Addr:    ":8080",
			Handler: e,
		},
	}
}

func (app *Application) Start() error {
	if err := app.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		app.logger.Println(err)
	}

	return nil
}

func (app *Application) Stop(error) {
	err := app.srv.Shutdown(app.ctx)
	if err != nil {
		app.logger.Println(err)
	}
}
