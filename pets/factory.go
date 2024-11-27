package pets

import (
	"github.com/yoshiMT2/breaders/models"
)

func NewPet(species string) *models.Pet {
	pet := models.Pet{
		Species:      species,
		Breed:        "",
		MinWeight:    0,
		MaxWeight:    0,
		Desicription: "no description entered yet",
		LifeSpan:     0,
	}

	return &pet
}
