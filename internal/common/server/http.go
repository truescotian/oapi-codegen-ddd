package server

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func RunHTTPServer(addr string, e *echo.Echo) {
	logrus.Info("Starting HTTP server")
	e.Logger.Fatal(e.Start(":1323"))
}
