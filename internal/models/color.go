package models

import (
	"equi_genea_horse_service/internal/db"
	horsepb "equi_genea_horse_service/internal/pb/api/horse"
)

type Color struct {
	ID          string
	Name        string
	Description *string
}

func (h *Color) LoadFromDB(dbHorseColor *db.HorseColor) *Color {
	h.ID = dbHorseColor.ID.String()
	h.Name = dbHorseColor.Name

	if dbHorseColor.Description.Valid {
		h.Description = &dbHorseColor.Description.String
	} else {
		h.Description = nil
	}

	return h
}

func (h *Color) ToHorseColorPB() *horsepb.HorseColor {
	return &horsepb.HorseColor{
		Id:          h.ID,
		Name:        h.Name,
		Description: h.Description,
	}
}
