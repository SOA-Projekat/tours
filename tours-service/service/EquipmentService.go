package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type EquipmentService struct {
	EquipmentRepo *repo.EquipmentRepository
}

func (service *EquipmentService) CreateEquipment(equipment *model.Equipment) error {
	// Attempt to create the equipment using the repository
	err := service.EquipmentRepo.CreateEquipment(equipment)
	if err != nil {
		return err
	}
	return nil
}

func (service *EquipmentService) GetEquipmentById(id string) (*model.Equipment, error) {
	equipment, err := service.EquipmentRepo.GetEquipmentById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("there are no equipment with id %s created", id))
	}
	return &equipment, nil
}

func (service *EquipmentService) UpdateEquipment(equipment *model.Equipment) error {
	err := service.EquipmentRepo.UpdateEquipment(equipment)
	if err != nil {
		return err
	}
	return nil
}

func (service *EquipmentService) DeleteEquipment(id int) error {
	return service.EquipmentRepo.DeleteEquipment(id)
}
