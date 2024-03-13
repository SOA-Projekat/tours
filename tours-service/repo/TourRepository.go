package repo

import (
	"database-example/model"
	"database/sql"
	"encoding/json"
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
	equipmentsJSONB, err := json.Marshal(tour.Equipments)
	if err != nil {
		return err
	}

	databaseResult := repository.DatabaseConnection.Model(&model.Tour{}).
		Where("id = ?", tour.ID).
		Update("equipments", equipmentsJSONB)

	if databaseResult.Error != nil {
		return databaseResult.Error
	}

	println("Rows affected: ", databaseResult.RowsAffected)
	return nil
}

func (repository *TourRepository) AddEquipmentToTour(tourID uint, equipmentIDs []uint) error {
	var tour model.Tour
	if err := repository.DatabaseConnection.First(&tour, tourID).Error; err != nil {
		return err
	}

	var equipments []model.Equipment
	if err := repository.DatabaseConnection.Find(&equipments, equipmentIDs).Error; err != nil {
		return err
	}

	tour.Equipments = equipments

	if err := repository.UpdateTour(&tour); err != nil {
		return err
	}

	return nil
}
func (repository *TourRepository) GetEquipmentForTour(tourID string) ([]model.Equipment, error) {
	var equipmentsJSON sql.NullString
	if err := repository.DatabaseConnection.Raw("SELECT equipments FROM tours WHERE id = ?", tourID).Scan(&equipmentsJSON).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch equipment data for tour with ID %s: %v", tourID, err)
	}

	if !equipmentsJSON.Valid {
		return nil, fmt.Errorf("equipment data for tour with ID %s is null", tourID)
	}

	var equipments []model.Equipment
	if err := json.Unmarshal([]byte(equipmentsJSON.String), &equipments); err != nil {
		return nil, fmt.Errorf("failed to unmarshal equipment data: %v", err)
	}

	return equipments, nil
}
