package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
	//"fmt"
)

type TourService struct {
	TourRepo      *repo.TourRepository
	EquipmentRepo *repo.EquipmentRepository
}

func (service *TourService) CreateTour(tour *model.Tour) error {
	err := service.TourRepo.CreateTour(tour)
	if err != nil {
		return err
	}
	return nil
}

func (service *TourService) GetTourById(id string) (*model.Tour, error) {
	tour, err := service.TourRepo.GetTourById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("tour with id %s is not found", id))
	}
	return &tour, nil
}

func (service *TourService) GetToursForAuthor(userId int) (*[]model.Tour, error) {
	tours, err := service.TourRepo.GetToursForAuthor(userId)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("there are no tours which user with id %d created", userId))
	}
	return &tours, nil
}

func (service *TourService) UpdateTour(tour *model.Tour) error {
	err := service.TourRepo.UpdateTour(tour)
	if err != nil {
		return err
	}
	return nil
}

func (service *TourService) AddEquipmentToTour(tourID uint, equipmentIDs []uint) error {
	err := service.TourRepo.AddEquipmentToTour(tourID, equipmentIDs)
	if err != nil {
		return err
	}
	return nil
}

func (service *TourService) GetEquipmentForTour(tourID string) ([]model.Equipment, error) {
	equipment, err := service.TourRepo.GetEquipmentForTour(tourID)
	if err != nil {
		return nil, err
	}
	return equipment, nil
}

func (service *TourService) GetAllEquipment() ([]model.Equipment, error) {
	allEquipment, err := service.EquipmentRepo.GetAllEquipment()
	if err != nil {
		return nil, err
	}
	return allEquipment, nil
}
