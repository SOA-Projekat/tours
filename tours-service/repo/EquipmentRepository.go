package repo

import (
	"database-example/model"
	"fmt"

	"gorm.io/gorm"
)

type EquipmentRepository struct {
	DatabaseConnection *gorm.DB
}

func (repository *EquipmentRepository) CreateEquipment(equipment *model.Equipment) error {
	var maximumId uint
	result := repository.DatabaseConnection.Model(&model.Equipment{}).Select("COALESCE(MAX(id),0)").Scan((&maximumId))
	if result.Error != nil {
		fmt.Println("couldnt find maximum id")
	}
	fmt.Printf("maximum id is %d\n", maximumId)

	equipment.ID = int(maximumId) + 1

	databaseCreationResult := repository.DatabaseConnection.Create(equipment)
	if databaseCreationResult != nil {
		return databaseCreationResult.Error
	}

	println("Rows affected: ", databaseCreationResult.RowsAffected)
	return nil
}

func (repository *TourRepository) GetEquipmentById(id string) (model.Equipment, error) {
	equipment := model.Equipment{}
	databaseResult := repository.DatabaseConnection.First(&equipment, "id = ?", id)
	if databaseResult != nil {
		return equipment, databaseResult.Error
	}
	return equipment, nil
}
