package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RunHTTPServer(addr string, e *echo.Echo) {
	// Start server
	go func() {
		if err := e.Start(":1323"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()
}
