package pet

import "context"

type Repository interface {
	GetPets(context.Context) []Pet
}
