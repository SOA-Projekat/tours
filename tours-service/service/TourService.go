package service

import (
	"database-example/model"
	"database-example/repo"
	"errors"
	"fmt"
	"time"
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

// func (service *TourService) PublishTour(tourID string) error {
// 	tour, err := service.TourRepo.GetTourById(tourID)
// 	if err != nil {
// 		return err
// 	}

// 	// Postavljanje vremena objavljivanja na trenutno vreme
// 	currentTime := time.Now()
// 	tour.PublishedDateTime = &currentTime

// 	// Ažuriranje ture u repozitorijumu
// 	err = service.TourRepo.UpdateTour(&tour)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func (service *TourService) PublishTour(tourID string) (*model.Tour, error) {
	//tour, err := service.TourRepo.GetTourById(tourID)
	tour, err := service.GetTourById(tourID)
	if err != nil {
		return nil, err
	}
	if tour.Name == "" {
		return nil, errors.New("tour must have a name to be published")
	}
	if tour.Description == "" {
		return nil, errors.New("tour must have a description to be published")
	}
	if tour.DifficultyLevel < 0 || tour.DifficultyLevel > 2 {
		return nil, errors.New("tour must have a valid difficulty level to be published")
	}
	if len(tour.Tags) == 0 {
		return nil, errors.New("tour must have at least one tag to be published")
	}
	if len(tour.TourPoints) < 2 {
		fmt.Println("Number of tour points: ", len(tour.TourPoints))
		return nil, errors.New("tour must have at least two key points to be published")
	}
	if tour.Status != 0 {
		return nil, errors.New("tour can be published only if its currently in draft phase")
	}
	// Postavljanje vremena arhiviranja na trenutno vreme
	currentTime := time.Now()
	tour.PublishedDateTime = currentTime
	tour.Status = 1

	// Ažuriranje ture u repozitorijumu
	err = service.TourRepo.UpdateTour(tour)
	if err != nil {
		return nil, err
	}

	return tour, nil
}

func (service *TourService) ArchiveTour(tourID string) (*model.Tour, error) {
	tour, err := service.GetTourById(tourID)
	if err != nil {
		return nil, err
	}
	if tour.Status == 0 || tour.Status != 1 {
		return nil, errors.New("tour can be tour can be archived only from published phase")
	}
	// Postavljanje vremena arhiviranja na trenutno vreme
	currentTime := time.Now()
	tour.ArchivedDateTime = currentTime
	tour.Status = 2

	// Ažuriranje ture u repozitorijumu
	err = service.TourRepo.UpdateTour(tour)
	if err != nil {
		return nil, err
	}

	return tour, nil

}

func (service *TourService) GetPublishedTours() ([]model.Tour, error) {
	// Retrieve published tours from the repository
	publishedTours, err := service.TourRepo.GetPublishedTours()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve published tours: %v", err)
	}

	// Iterate through the published tours and fetch associated equipments and tour points
	for i := range publishedTours {
		// Fetch and assign equipments to each tour
		equipments, err := service.EquipmentRepo.GetEquipmentByTourID(publishedTours[i].ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get equipments for tour %d: %v", publishedTours[i].ID, err)
		}
		publishedTours[i].Equipments = equipments

		// Fetch and assign tour points to each tour
		tourPoints, err := service.TourPointRepo.GetTourPointsByTourId(publishedTours[i].ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get tour points for tour %d: %v", publishedTours[i].ID, err)
		}
		publishedTours[i].TourPoints = tourPoints
	}

	return publishedTours, nil
}
