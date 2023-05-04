package routes

import (
	"am/office-check-in/middlewares"
	"am/office-check-in/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddEmployees(e *echo.Echo) {
	g := e.Group("/employees")

	g.Use(middlewares.AuthWithToken)
	g.GET("", getAll)
	g.GET("/logs/:id", getLogs)
}

func getAll(c echo.Context) error {
	users, err := services.GetAllEmployees()

	if err != nil {
		return c.String(http.StatusBadRequest, "could not find any users")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"employees": users,
	})
}

func getLogs(c echo.Context) error {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)

	if err != nil {
		return c.String(http.StatusBadRequest, "id is not integer")
	}

	logs, err := services.GetLogsForEmployee(uint(intId))

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"logs": logs,
	})
}
