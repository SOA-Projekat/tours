package service

import (
	"database-example/model"
	"database-example/repo"
	//"fmt"
)

type TourService struct {
	TourRepo *repo.TourRepository
}

func (service *TourService) CreateTour(tour *model.Tour) error {
	err := service.TourRepo.CreateTour(tour)
	if err != nil {
		return err
	}
	return nil
}
