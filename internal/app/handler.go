package app

import (
	"context"
	horsepb "equi_genea_horse_service/internal/pb/api/horse"
	"equi_genea_horse_service/internal/service"

	"google.golang.org/protobuf/types/known/emptypb"
)

type HorseHandler struct {
	service *service.HorseService
	horsepb.UnimplementedHorseServiceServer
}

func NewHorseHandler(service *service.HorseService) *HorseHandler {
	return &HorseHandler{service: service}
}

func (h *HorseHandler) GetGeneticMarkerList(ctx context.Context, _ *emptypb.Empty) (*horsepb.GetGeneticMarkerListResponse, error) {
	markers, err := h.service.GetGeneticMarkerList(ctx)
	if err != nil {
		return nil, err
	}

	pbMarkers := make([]*horsepb.HorseGeneticMarker, len(markers))
	for i := 0; i < len(markers); i++ {
		pbMarkers[i] = markers[i].ToHorseGeneticMarkerPB()
	}

	return &horsepb.GetGeneticMarkerListResponse{Markers: pbMarkers}, nil
}

func (h *HorseHandler) GetBirthplaceList(ctx context.Context, _ *emptypb.Empty) (*horsepb.GetBirthplaceListResponse, error) {
	birthplaces, err := h.service.GetBirthplaceList(ctx)
	if err != nil {
		return nil, err
	}

	pbBirthplaces := make([]*horsepb.HorseBirthplace, len(birthplaces))
	for i := 0; i < len(birthplaces); i++ {
		pbBirthplaces[i] = birthplaces[i].ToHorseBirtplacePB()
	}

	return &horsepb.GetBirthplaceListResponse{Birthplaces: pbBirthplaces}, nil
}

func (h *HorseHandler) GetColorList(ctx context.Context, _ *emptypb.Empty) (*horsepb.GetColorListResponse, error) {
	colors, err := h.service.GetColorList(ctx)
	if err != nil {
		return nil, err
	}

	pbColors := make([]*horsepb.HorseColor, len(colors))
	for i := 0; i < len(colors); i++ {
		pbColors[i] = colors[i].ToHorseColorPB()
	}

	return &horsepb.GetColorListResponse{Colors: pbColors}, nil
}

func (h *HorseHandler) GetGenderList(ctx context.Context, _ *emptypb.Empty) (*horsepb.GetGenderListResponse, error) {
	genders, err := h.service.GetGenderList(ctx)
	if err != nil {
		return nil, err
	}

	pbGenders := make([]*horsepb.HorseGender, len(genders))
	for i := 0; i < len(genders); i++ {
		pbGenders[i] = genders[i].ToHorseGenderPB()
	}

	return &horsepb.GetGenderListResponse{Genders: pbGenders}, nil
}

func (h *HorseHandler) CreateGender(ctx context.Context, in *horsepb.CreateGenderRequest) (*horsepb.CreateGenderResponse, error) {
	gender, err := h.service.CreateGender(ctx, in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	return &horsepb.CreateGenderResponse{HorseGender: gender.ToHorseGenderPB()}, nil
}
