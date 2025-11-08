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

func (s *HorseService) CreateGender(ctx context.Context, name string, description *string) (*models.Gender, error) {
	dbGender, err := s.queries.CreateHorseGender(ctx, db.CreateHorseGenderParams{
		Name:        name,
		Description: *pgutil.StringPtrToPgText(description),
	})
	if err != nil {
		return nil, err
	}

	gender := models.Gender{}

	return gender.LoadFromDB(&dbGender), nil
}

func (s *HorseService) GetGenderList(ctx context.Context) ([]*models.Gender, error) {
	dbGenders, err := s.queries.GetHorseGenderList(ctx)
	if err != nil {
		return nil, err
	}

	genders := make([]*models.Gender, len(dbGenders))

	for i := 0; i < len(dbGenders); i++ {
		genders[i] = &models.Gender{}
		genders[i].LoadFromDB(&dbGenders[i])
	}

	return genders, nil
}

func (s *HorseService) GetColorList(ctx context.Context) ([]*models.Color, error) {
	dbColors, err := s.queries.GetHorseColorList(ctx)
	if err != nil {
		return nil, err
	}

	colors := make([]*models.Color, len(dbColors))

	for i := 0; i < len(dbColors); i++ {
		colors[i] = &models.Color{}
		colors[i].LoadFromDB(&dbColors[i])
	}

	return colors, nil
}

func (s *HorseService) GetBirthplaceList(ctx context.Context) ([]*models.Birthplace, error) {
	dbBirthplaces, err := s.queries.GetHorseBirthplaceList(ctx)
	if err != nil {
		return nil, err
	}

	birthplaces := make([]*models.Birthplace, len(dbBirthplaces))

	for i := 0; i < len(dbBirthplaces); i++ {
		birthplaces[i] = &models.Birthplace{}
		birthplaces[i].LoadFromDB(&dbBirthplaces[i])
	}

	return birthplaces, nil
}

func (s *HorseService) GetGeneticMarkerList(ctx context.Context) ([]*models.GeneticMarker, error) {
	dbGeneticMarkers, err := s.queries.GetHorseGeneticMarkerList(ctx)
	if err != nil {
		return nil, err
	}

	markers := make([]*models.GeneticMarker, len(dbGeneticMarkers))

	for i := 0; i < len(dbGeneticMarkers); i++ {
		markers[i] = &models.GeneticMarker{}
		markers[i].LoadFromDB(&dbGeneticMarkers[i])
	}

	return markers, nil
}
