package adapters

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	domain "github.com/truescotian/oapi-codegen-example/internal/petstore/domain/pet"
)

type PetModel struct {
	Name string
}

type PostgresPetRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresPetRepository(pool *pgxpool.Pool) PostgresPetRepository {
	return PostgresPetRepository{pool: pool}
}

func (p PostgresPetRepository) GetPets(ctx context.Context) []domain.Pet {
	return make([]domain.Pet, 0)
}
