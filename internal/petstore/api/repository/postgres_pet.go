package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/truescotian/oapi-codegen-example/internal/petstore/api/domain"
	"github.com/truescotian/oapi-codegen-example/internal/petstore/api/models"
)

type PostgresPetRepository struct {
	conn Connection
}

func NewPostgresPetRepository(conn Connection) domain.PetRepository {
	return PostgresPetRepository{conn: conn}
}

func (p PostgresPetRepository) GetPets(ctx echo.Context) []models.Pet {
	return make([]models.Pet, 0)
}
