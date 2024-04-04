package repo

import (
	"database-example/model"
	"fmt"

	"gorm.io/gorm"
)

type TourRepository struct {
	DatabaseConnection *gorm.DB
}

func (repository *TourRepository) CreateTour(tour *model.Tour) error {
	var maximumId uint
	res := repository.DatabaseConnection.Model(&model.Tour{}).Select("COALESCE(MAX(id),0)").Scan((&maximumId))
	if res.Error != nil {
		fmt.Println("coundn't find maximum ID")
	}
	fmt.Printf("Maximum id is %d\n", maximumId)

	tour.ID = int(maximumId) + 1
	fmt.Println("tour:", tour)

	dbCreationResult := repository.DatabaseConnection.Create(tour)
	if dbCreationResult != nil {
		return dbCreationResult.Error
	}

	println("Rows affected: ", dbCreationResult.RowsAffected)
	return nil
}

func (repository *TourRepository) GetTourById(id string) (model.Tour, error) {
	tour := model.Tour{}
	databaseResult := repository.DatabaseConnection.First(&tour, "id = ?", id)
	if databaseResult != nil {
		return tour, databaseResult.Error
	}
	return tour, nil
}

func (repository *TourRepository) GetToursForAuthor(userId int) ([]model.Tour, error) {
	tours := []model.Tour{}
	databaseResult := repository.DatabaseConnection.Where("user_id = ?", userId).Find(&tours)
	if databaseResult != nil {
		return tours, databaseResult.Error
	}
	return tours, nil
}

func (repository *TourRepository) UpdateTour(tour *model.Tour) error {
	databaseResult := repository.DatabaseConnection.Save(tour)

	if databaseResult.Error != nil {
		return databaseResult.Error
	}
	println("Rows affected: ", databaseResult.RowsAffected)
	return nil
}

func (repository *TourRepository) AddEquipmentToTour(tourID int, equipmentID int) error {
	// Retrieve the tour by its ID
	var tour model.Tour
	if err := repository.DatabaseConnection.First(&tour, tourID).Error; err != nil {
		return fmt.Errorf("failed to find tour with ID %d: %w", tourID, err)
	}

	// Retrieve the equipment by its ID
	var equipment model.Equipment
	if err := repository.DatabaseConnection.First(&equipment, equipmentID).Error; err != nil {
		return fmt.Errorf("failed to find equipment with ID %d: %w", equipmentID, err)
	}

	// Set the TourID of the equipment to the provided tourID
	equipment.TourID = tourID

	// Add the equipment to the tour's equipments slice
	tour.Equipments = append(tour.Equipments, equipment)

	if err := repository.DatabaseConnection.Save(&tour).Error; err != nil {
		return fmt.Errorf("failed to update tour with ID %d: %w", tourID, err)
	}

	return nil
}

func (repository *TourRepository) GetPublishedTours() ([]model.Tour, error) {
	var tours []model.Tour
	databaseResult := repository.DatabaseConnection.Where("status = ?", 1).Find(&tours)
	if databaseResult.Error != nil {
		return nil, databaseResult.Error
	}
	return tours, nil
}
