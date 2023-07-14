package domain

import (
	"github.com/labstack/echo/v4"
	"github.com/truescotian/oapi-codegen-example/internal/petstore/api/models"
)

type PetRepository interface {
	GetPets(echo.Context) []models.Pet
}
