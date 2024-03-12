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
