package repo

import (
	"database-example/model"
	"fmt"

	"gorm.io/gorm"
)

type EquipmentRepository struct {
	DatabaseConnection *gorm.DB
}

// EquipmentRepository
func (repository *EquipmentRepository) CreateEquipment(equipment *model.Equipment) error {
	// Check if the referenced tour exists before creating the equipment
	var count int64
	repository.DatabaseConnection.Model(&model.Tour{}).Where("id = ?", equipment.TourID).Count(&count)
	if count == 0 {
		return fmt.Errorf("failed to create equipment: tour with ID %d does not exist", equipment.TourID)
	}

	// Attempt to create the equipment in the database
	databaseCreationResult := repository.DatabaseConnection.Create(equipment)
	if databaseCreationResult.Error != nil {
		return fmt.Errorf("failed to create equipment: %v", databaseCreationResult.Error)
	}

	fmt.Printf("Equipment created with ID: %d\n", equipment.ID)
	return nil
}

func (repository *EquipmentRepository) GetEquipmentById(id string) (model.Equipment, error) {
	equipment := model.Equipment{}
	databaseResult := repository.DatabaseConnection.First(&equipment, "id = ?", id)
	if databaseResult != nil {
		return equipment, databaseResult.Error
	}
	return equipment, nil
}

func (repository *EquipmentRepository) UpdateEquipment(equipment *model.Equipment) error {
	databaseResult := repository.DatabaseConnection.Save(equipment)

	if databaseResult.Error != nil {
		return databaseResult.Error
	}
	println("Rows affected: ", databaseResult.RowsAffected)
	return nil
}

func (repository *EquipmentRepository) DeleteEquipment(id int) error {
	result := repository.DatabaseConnection.Delete(&model.Equipment{}, id)
	if result.Error != nil {
		return result.Error
	}
	println("Rows affected: ", result.RowsAffected)
	return nil
}

func (repository *EquipmentRepository) GetEquipmentByTourID(tourID int) ([]model.Equipment, error) {
	var equipments []model.Equipment
	if err := repository.DatabaseConnection.Find(&equipments, "tour_id = ?", tourID).Error; err != nil {
		return nil, err
	}
	return equipments, nil
}
