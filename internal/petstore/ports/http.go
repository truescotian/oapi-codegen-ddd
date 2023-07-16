package ports

import (
	"net/http"

	"github.com/deepmap/oapi-codegen/examples/petstore-expanded/echo/api/models"
	"github.com/labstack/echo/v4"
	"github.com/truescotian/oapi-codegen-example/internal/petstore/domain/pet"
)

type HttpServer struct {
	petRepo pet.Repository
}

func NewHttpServer(pr pet.Repository) HttpServer {
	return HttpServer{petRepo: pr}
}

func (h *HttpServer) FindPets(ctx echo.Context, request FindPetsParams) error {
	pets := h.petRepo.GetPets(ctx.Request().Context())

	result := make([]models.Pet, 0)

	for _, pet := range pets {
		result = append(result, models.Pet{
			Id:   1,
			Name: pet.Name,
			Tag:  nil,
		})
	}

	return ctx.JSON(http.StatusOK, result)
}
