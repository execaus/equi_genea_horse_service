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

func (h *HorseHandler) CreateGender(ctx context.Context, in *horsepb.CreateGenderRequest) (*horsepb.CreateGenderResponse, error) {
	gender, err := h.service.CreateGender(ctx, in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	return &horsepb.CreateGenderResponse{HorseGender: gender.ToHorseGenderPB()}, nil
}

func (h *HorseHandler) GetGenderList(ctx context.Context, in *emptypb.Empty) (*horsepb.GetGenderListResponse, error) {
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
