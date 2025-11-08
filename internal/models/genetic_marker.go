package models

import (
	"equi_genea_horse_service/internal/db"
	horsepb "equi_genea_horse_service/internal/pb/api/horse"
)

type GeneticMarker struct {
	ID          string
	Name        string
	Description *string
}

func (h *GeneticMarker) LoadFromDB(dbGeneticMarker *db.GeneticMarker) *GeneticMarker {
	h.ID = dbGeneticMarker.ID.String()
	h.Name = dbGeneticMarker.Name

	if dbGeneticMarker.Description.Valid {
		h.Description = &dbGeneticMarker.Description.String
	} else {
		h.Description = nil
	}

	return h
}

func (h *GeneticMarker) ToHorseGeneticMarkerPB() *horsepb.HorseGeneticMarker {
	return &horsepb.HorseGeneticMarker{
		Id:          h.ID,
		Name:        h.Name,
		Description: h.Description,
	}
}
