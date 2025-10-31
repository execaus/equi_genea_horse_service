package service

import (
	"context"
	"equi_genea_horse_service/internal/db"
	"equi_genea_horse_service/internal/models"
	"equi_genea_horse_service/internal/pgutil"
)

type HorseService struct {
	queries *db.Queries
}

func NewHorseService(queries *db.Queries) *HorseService {
	return &HorseService{queries: queries}
}

func (h *HorseService) CreateGender(ctx context.Context, name string, description *string) (*models.HorseGender, error) {
	dbGender, err := h.queries.CreateHorseGender(ctx, db.CreateHorseGenderParams{
		Name:        name,
		Description: *pgutil.StringPtrToPgText(description),
	})
	if err != nil {
		return nil, err
	}

	gender := models.HorseGender{}

	return gender.LoadFromDB(&dbGender), nil
}

func (h *HorseService) GetGenderList() []*models.HorseGender {

}
