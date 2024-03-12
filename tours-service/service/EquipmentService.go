package service

import (
	"database-example/model"
	"database-example/repo"
)

type EquipmentService struct {
	EquipmentRepo *repo.EquipmentRepository
}

func (service *EquipmentService) CreateEquipment(equipment *model.Equipment) error {
	err := service.EquipmentRepo.CreateEquipment(equipment)
	if err != nil {
		return err
	}
	return nil
}
