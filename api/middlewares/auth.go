package middlewares

import (
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// ** An implementation of echo's KeyAuthWithConfig that parses a given token and continues if valid **
var AuthWithToken = middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
	Validator: func(auth string, c echo.Context) (bool, error) {
		// Try parsing given token
		_, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) { return []byte(os.Getenv("JWT_SECRET")), nil })

		// If error thrown, token was invalid
		if err != nil {
			return false, err
		}

		return true, nil
	},
})
