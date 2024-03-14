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
