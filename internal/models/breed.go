package models

import (
	"equi_genea_horse_service/internal/db"
	horsepb "equi_genea_horse_service/internal/pb/api/horse"
)

type Breed struct {
	ID          string
	Name        string
	Description *string
}

func (h *Breed) LoadFromDB(dbHorseBreed *db.HorseBreed) *Breed {
	h.ID = dbHorseBreed.ID.String()
	h.Name = dbHorseBreed.Name

	if dbHorseBreed.Description.Valid {
		h.Description = &dbHorseBreed.Description.String
	} else {
		h.Description = nil
	}

	return h
}

func (h *Breed) ToHorseBreedPB() *horsepb.HorseBreed {
	return &horsepb.HorseBreed{
		Id:          h.ID,
		Name:        h.Name,
		Description: h.Description,
	}
}
