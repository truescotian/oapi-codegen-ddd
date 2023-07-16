package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/labstack/echo/v4"
	"github.com/truescotian/oapi-codegen-example/internal/cmdutil"
	"github.com/truescotian/oapi-codegen-example/internal/common/server"
	"github.com/truescotian/oapi-codegen-example/internal/petstore/api"
	"github.com/truescotian/oapi-codegen-example/internal/petstore/api/models"
	"github.com/truescotian/oapi-codegen-example/internal/petstore/api/repository"
)

func Execute(ctx context.Context) int {
	db, err := cmdutil.NewDatabasePool(ctx, 1)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connected")
	}

	petRepo := repository.NewPostgresPetRepository(db)

	myApi := api.HttpServer{PetRepo: petRepo}

	e := echo.New()

	models.RegisterHandlers(e, &myApi)

	server.RunHTTPServer(":3000", e)

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	return 1
}
