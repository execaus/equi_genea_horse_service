package service

import "equi_genea_horse_service/internal/db"

type HorseService struct {
	queries *db.Queries
}

func NewHorseService(queries *db.Queries) *HorseService {
	return &HorseService{queries: queries}
}
