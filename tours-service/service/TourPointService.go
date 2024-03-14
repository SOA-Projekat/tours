package service

import (
	"database-example/model"
	"database-example/repo"
	//"fmt"
)

type TourPointService struct {
	TourPointRepo *repo.TourPointRepository
}

func (service *TourPointService) Create(tourPoint *model.TourPoint) error {
	err := service.TourPointRepo.CreateTourPoint(tourPoint)
	if err != nil {
		return err
	}
	return nil
}

func (service *TourPointService) GetTourPointsByTourId(tourId int) ([]model.TourPoint, error) {
	return service.TourPointRepo.GetTourPointsByTourId(tourId)
}

func (service *TourPointService) UpdateTourPoint(tourPoint *model.TourPoint) error {
	err := service.TourPointRepo.UpdateTourPoint(tourPoint)
	if err != nil {
		return err
	}
	return nil
}

func (service *TourPointService) DeleteTourPoint(id int) error {
	return service.TourPointRepo.DeleteTourPoint(id)
}
