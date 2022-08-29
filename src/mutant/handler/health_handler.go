package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github/meli/src/mutant/usecase"
)

type mutanHandler struct {
	useCase usecase.MutantUsecase
}

func NewMutantHandler(e *echo.Echo, useCase usecase.MutantUsecase) {
	h := &mutanHandler{useCase}
	e.POST("/mutant", h.IsMutant)
	e.GET("/mutant/stats", h.GetStats)
}

func (h *mutanHandler) GetStats(c echo.Context) error {

	rst, err := h.useCase.GetStats()
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{"error": err})
	}
	return c.JSON(http.StatusOK, rst)
}

func (h *mutanHandler) IsMutant(c echo.Context) error {
	var ada []string
	rst, err := h.useCase.IsMutant(ada)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if rst.IsMutant {
		return c.JSON(http.StatusOK, rst)
	}

	return c.JSON(http.StatusForbidden, rst)
}
