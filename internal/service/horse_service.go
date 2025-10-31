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

func (h *HorseService) GetGenderList(ctx context.Context) ([]*models.HorseGender, error) {
	dbGenders, err := h.queries.GetHorseGenderList(ctx)
	if err != nil {
		return nil, err
	}

	genders := make([]*models.HorseGender, len(dbGenders))

	var currGender *models.HorseGender

	for i := 0; i < len(dbGenders); i++ {
		currGender = &models.HorseGender{}
		genders = append(genders, currGender.LoadFromDB(&dbGenders[i]))
	}

	return genders, nil
}
