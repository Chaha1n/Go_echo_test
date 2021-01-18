package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Hello World!")
	})
	e.Logger.Fatal(e.Start(":1234"))
}
