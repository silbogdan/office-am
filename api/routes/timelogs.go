package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"am/office-check-in/models"
	"am/office-check-in/services"
)

func AddTimeLogs(e *echo.Echo) {
	g := e.Group("/timelogs")

	g.POST("/scan", scan)
}

func scan(c echo.Context) error {
	tl := new(models.TimeLogBody)

	// Get entrance type
	if err := c.Bind(tl); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	// Scan QR code (using camera)
	scanResult, err := services.Scan(tl.Type)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"qrCodeData": scanResult,
	})
}
