package models

import (
	"equi_genea_horse_service/internal/db"
	horsepb "equi_genea_horse_service/internal/pb/api/horse"
)

type HorseGender struct {
	ID          string
	Name        string
	Description *string
}

func (h *HorseGender) LoadFromDB(dbHorseGender *db.HorseGender) *HorseGender {
	h.ID = dbHorseGender.ID.String()
	h.Name = dbHorseGender.Name

	if dbHorseGender.Description.Valid {
		h.Description = &dbHorseGender.Description.String
	} else {
		h.Description = nil
	}

	return h
}

func (h *HorseGender) ToHorseGenderPB() *horsepb.HorseGender {
	return &horsepb.HorseGender{
		Id:          h.ID,
		Name:        h.Name,
		Description: h.Description,
	}
}
