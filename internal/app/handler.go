package app

import (
	horsepb "equi_genea_horse_service/internal/pb/api/horse"
	"equi_genea_horse_service/internal/service"
)

type HorseHandler struct {
	service *service.HorseService
	horsepb.UnimplementedHorseServiceServer
}

func NewHorseHandler(service *service.HorseService) *HorseHandler {
	return &HorseHandler{service: service}
}
