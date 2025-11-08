package models

import (
	"equi_genea_horse_service/internal/db"
	horsepb "equi_genea_horse_service/internal/pb/api/horse"
)

type HorseBirthplace struct {
	ID          string
	Name        string
	Description *string
}

func (h *HorseBirthplace) LoadFromDB(dbHorseBirthplace *db.HorseBirthplace) *HorseBirthplace {
	h.ID = dbHorseBirthplace.ID.String()
	h.Name = dbHorseBirthplace.Name

	if dbHorseBirthplace.Description.Valid {
		h.Description = &dbHorseBirthplace.Description.String
	} else {
		h.Description = nil
	}

	return h
}

func (h *HorseBirthplace) ToHorseColorPB() *horsepb.HorseBirthplace {
	return &horsepb.HorseBirthplace{
		Id:          h.ID,
		Name:        h.Name,
		Description: h.Description,
	}
}
