package models

import (
	"equi_genea_horse_service/internal/db"
	horsepb "equi_genea_horse_service/internal/pb/api/horse"
)

type Birthplace struct {
	ID          string
	Name        string
	Description *string
}

func (h *Birthplace) LoadFromDB(dbHorseBirthplace *db.HorseBirthplace) *Birthplace {
	h.ID = dbHorseBirthplace.ID.String()
	h.Name = dbHorseBirthplace.Name

	if dbHorseBirthplace.Description.Valid {
		h.Description = &dbHorseBirthplace.Description.String
	} else {
		h.Description = nil
	}

	return h
}

func (h *Birthplace) ToHorseBirtplacePB() *horsepb.HorseBirthplace {
	return &horsepb.HorseBirthplace{
		Id:          h.ID,
		Name:        h.Name,
		Description: h.Description,
	}
}
