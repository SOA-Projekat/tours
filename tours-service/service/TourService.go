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
	TourPointRepo *repo.TourPointRepository
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
	equipments, err := service.EquipmentRepo.GetEquipmentByTourID(tour.ID)
	if err != nil {
		return nil, err
	}

	tour.Equipments = equipments

	tourPoints, err := service.TourPointRepo.GetTourPointsByTourId(tour.ID)
	if err != nil {
		return nil, err
	}
	tour.TourPoints = tourPoints
	return &tour, nil
}

func (service *TourService) GetToursForAuthor(userId int) ([]model.Tour, error) {
	tours, err := service.TourRepo.GetToursForAuthor(userId)
	if err != nil {
		return nil, fmt.Errorf("there are no tours which user with id %d created: %v", userId, err)
	}

	for i := range tours {
		// Fetch and assign equipments to each tour
		equipments, err := service.EquipmentRepo.GetEquipmentByTourID(tours[i].ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get equipments for tour %d: %v", tours[i].ID, err)
		}
		tours[i].Equipments = equipments

		// Fetch and assign tour points to each tour
		tourPoints, err := service.TourPointRepo.GetTourPointsByTourId(tours[i].ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get tour points for tour %d: %v", tours[i].ID, err)
		}
		tours[i].TourPoints = tourPoints
	}

	return tours, nil
}

/*
	func (service *TourService) GetToursForAuthor(userId int) (*[]model.Tour, error) {
		tours, err := service.TourRepo.GetToursForAuthor(userId)
		if err != nil {
			return nil, fmt.Errorf(fmt.Sprintf("there are no tours which user with id %d created", userId))
		}
		return &tours, nil
	}
*/
func (service *TourService) UpdateTour(tour *model.Tour) error {
	err := service.TourRepo.UpdateTour(tour)
	if err != nil {
		return err
	}
	return nil
}

func (service *TourService) AddEquipmentToTour(tourID int, equipmentID int) error {
	err := service.TourRepo.AddEquipmentToTour(tourID, equipmentID)
	if err != nil {
		return err
	}
	return nil
}
