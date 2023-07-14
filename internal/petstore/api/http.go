package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/truescotian/oapi-codegen-example/internal/petstore/api/domain"
	"github.com/truescotian/oapi-codegen-example/internal/petstore/api/models"
)

type HttpServer struct {
	PetRepo domain.PetRepository
}

func (h *HttpServer) FindPets(ctx echo.Context, request models.FindPetsParams) error {
	pets := h.PetRepo.GetPets(ctx)

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
