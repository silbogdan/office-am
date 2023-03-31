package routes

import (
	"am/office-check-in/services"
	"fmt"

	"github.com/labstack/echo/v4"
)

func AddFile(e *echo.Echo) {
	g := e.Group("/file")

	g.POST("/upload", upload)
}

func upload(c echo.Context) error {
	file, err := c.FormFile("file")

	if err != nil {
		return c.String(500, "failed to get file")
	}

	src, err := file.Open()

	if err != nil {
		return c.String(500, "failed to open file")
	}

	defer src.Close()

	url, err := services.Upload(src, file.Filename, file.Header.Get("Content-Type"))

	if err != nil {
		fmt.Println(err)
		return c.String(500, "failed to upload file")
	}

	return c.JSON(200, map[string]string{
		"url": url,
	})
}
