package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"am/office-check-in/models"
	"am/office-check-in/services"
)

func AddAuth(e *echo.Echo) {
	g := e.Group("/auth")

	g.POST("/register", register)
	g.POST("/login", login)
}

func register(c echo.Context) error {
	u := new(models.UserBody)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	user := models.UserBody{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Picture:  u.Picture,
	}

	createdUser, err := services.Register(user)

	if err != nil {
		return c.String(http.StatusBadRequest, "failed to create user")
	}

	return c.JSON(http.StatusOK, createdUser)
}

func login(c echo.Context) error {
	u := new(models.UserLoginBody)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	token, err := services.Login(u.Email, u.Password)

	if err != nil {
		return c.String(http.StatusBadRequest, "failed to create user")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}
