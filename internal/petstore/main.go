package petstore

import (
	"context"
	"os"
	"os/signal"

	"github.com/labstack/echo/v4"
	"github.com/truescotian/oapi-codegen-example/internal/cmdutil"
	"github.com/truescotian/oapi-codegen-example/internal/common/server"
	"github.com/truescotian/oapi-codegen-example/internal/petstore/api"
	"github.com/truescotian/oapi-codegen-example/internal/petstore/api/models"
	"github.com/truescotian/oapi-codegen-example/internal/petstore/api/repository"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	db, err := cmdutil.NewDatabasePool(ctx, 1)
	if err != nil {
		panic(err)
	}

	petRepo := repository.NewPostgresPetRepository(db)

	myApi := api.HttpServer{PetRepo: petRepo}

	e := echo.New()

	models.RegisterHandlers(e, &myApi)

	server.RunHTTPServer(":3000", e)
}
