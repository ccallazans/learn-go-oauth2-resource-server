package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type healthResponse struct {
	Status string `json:"status"`
}

func Health(c echo.Context) error {
	status := &healthResponse{
		Status: "OK",
	}

	return c.JSON(http.StatusOK, status)
}
