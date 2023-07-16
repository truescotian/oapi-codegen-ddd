package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/truescotian/oapi-codegen-example/internal/common/server"
	"github.com/truescotian/oapi-codegen-example/internal/petstore/ports"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	app := NewApplication(ctx)

	myApi := ports.NewHttpServer(app.petRepo)

	e := echo.New()

	ports.RegisterHandlers(e, &myApi)

	server.RunHTTPServer(":3000", e)

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	os.Exit(1)
}
