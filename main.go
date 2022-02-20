package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", handleRoot)
	port := os.Getenv("PORT")
	listen := fmt.Sprintf(":%s", port)
	e.Logger.Fatal(e.Start(listen))
}

func handleRoot(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
