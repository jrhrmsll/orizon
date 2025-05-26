package http

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/jrhrmsll/orizon"
)

type IntervalController struct {
	intervalService orizon.IntervalService
}

func NewIntervalController(intervalService orizon.IntervalService) *IntervalController {
	return &IntervalController{
		intervalService: intervalService,
	}
}

func (controller *IntervalController) Index(c echo.Context) error {
	payload := new(orizon.IntervalSpec)
	if err := c.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, controller.intervalService.Find(payload))
}
