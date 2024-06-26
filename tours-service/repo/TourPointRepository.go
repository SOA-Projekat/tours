package repo

import (
	"database-example/model"
	"fmt"

	"gorm.io/gorm"
)

type TourPointRepository struct {
	DatabaseConnection *gorm.DB
}

func (keypointRepository *TourPointRepository) CreateTourPoint(tourPoint *model.TourPoint) error {
	var maxId uint
	res := keypointRepository.DatabaseConnection.Model(&model.TourPoint{}).Select("COALESCE(MAX(id),0)").Scan((&maxId))
	if res.Error != nil {
		fmt.Println("cant fint max id")
	}
	fmt.Printf("max id is %d\n", maxId)

	tourPoint.ID = int(maxId) + 1
	fmt.Println("tour point: ", tourPoint)

	dbResult := keypointRepository.DatabaseConnection.Create(tourPoint)
	if dbResult != nil {
		return dbResult.Error
	}
	return nil
}
func (repo *TourPointRepository) GetTourPointsByTourId(tourId int) ([]model.TourPoint, error) {
	var tourPoints []model.TourPoint
	if err := repo.DatabaseConnection.Where("tour_id = ?", tourId).Find(&tourPoints).Error; err != nil {
		return nil, err
	}
	return tourPoints, nil
}

func (repo *TourPointRepository) UpdateTourPoint(tourPoint *model.TourPoint) error {
	databaseResult := repo.DatabaseConnection.Save(tourPoint)

	if databaseResult.Error != nil {
		return databaseResult.Error
	}
	println("Rows affected: ", databaseResult.RowsAffected)
	return nil
}

func (repository *TourPointRepository) DeleteTourPoint(id int) error {
	result := repository.DatabaseConnection.Delete(&model.TourPoint{}, id)
	if result.Error != nil {
		return result.Error
	}
	println("Rows affected: ", result.RowsAffected)
	return nil
}
