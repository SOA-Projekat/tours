package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"log"
	"strconv"

	//"fmt"
	//"log"
	"net/http"

	"github.com/gorilla/mux"
	//"strconv"
	//"github.com/gorilla/mux"
)

type TourHandler struct {
	TourService *service.TourService
}

func (handler *TourHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var tour model.Tour
	err := json.NewDecoder(req.Body).Decode(&tour)
	if err != nil {
		println("error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.TourService.CreateTour(&tour)
	if err != nil {
		println("error while creating tour")
		writer.WriteHeader(http.StatusExpectationFailed)
		return

	}
	writer.WriteHeader((http.StatusCreated))
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(tour)
}

func (handler *TourHandler) GetTourById(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Tour with requested id %s", id)
	tour, err := handler.TourService.GetTourById(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(tour)
}

// func (handler *TourHandler) GetToursForAuthor(writer http.ResponseWriter, req *http.Request) {
// 	id := mux.Vars(req)["userId"]
// 	converterId, err := strconv.Atoi(id)
// 	if err != nil {
// 		fmt.Println("integer can't nbe converted to integer")
// 	}
// 	tours, err := handler.TourService.GetToursForAuthor(converterId)
// 	writer.Header().Set("Content-Type", "application/json")
// 	if err != nil {
// 		writer.WriteHeader(http.StatusNotFound)
// 		return
// 	}
// 	writer.WriteHeader(http.StatusOK)
// 	json.NewEncoder(writer).Encode(tours)
// }

func (handler *TourHandler) UpdateTour(writer http.ResponseWriter, req *http.Request) {
	var tour model.Tour

	err := json.NewDecoder(req.Body).Decode(&tour)
	if err != nil {
		println("error occured whiule parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.TourService.UpdateTour(&tour)
	if err != nil {
		println("error occured while updating tour")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(tour)

}

func (handler *TourHandler) AddEquipmentToTour(writer http.ResponseWriter, req *http.Request) {
	// Extract tour ID and equipment ID from the request parameters
	tourID := mux.Vars(req)["tourID"]
	equipmentID := mux.Vars(req)["equipmentID"]

	// Convert IDs to integers
	tourIDInt, err := strconv.Atoi(tourID)
	if err != nil {
		log.Printf("Invalid tour ID: %s", tourID)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	equipmentIDInt, err := strconv.Atoi(equipmentID)
	if err != nil {
		log.Printf("Invalid equipment ID: %s", equipmentID)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// Call the TourService method to add equipment to the tour
	err = handler.TourService.AddEquipmentToTour(tourIDInt, equipmentIDInt)
	if err != nil {
		log.Printf("Error adding equipment to tour: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Respond with success status
	writer.WriteHeader(http.StatusOK)
}

func (handler *TourHandler) PublishTour(writer http.ResponseWriter, req *http.Request) {
	tourID := mux.Vars(req)["tourID"]
	tour, err := handler.TourService.PublishTour(tourID)
	if err != nil {
		log.Printf("Error while publishing %v", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(tour)
}

func (handler *TourHandler) ArchiveTour(writer http.ResponseWriter, req *http.Request) {
	tourID := mux.Vars(req)["tourID"]
	tour, err := handler.TourService.ArchiveTour(tourID)
	if err != nil {
		log.Printf("Error while archiving: %v", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(tour)
}

func (handler *TourHandler) GetPublishedTours(writer http.ResponseWriter, req *http.Request) {
	// Call the GetPublishedTours method from TourService
	publishedTours, err := handler.TourService.GetPublishedTours()
	if err != nil {
		log.Printf("Error retrieving published tours: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Respond with the fetched published tours
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(publishedTours); err != nil {
		log.Printf("Error encoding published tours: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

//TODO: probably keypoint logic is added here i suppose. ask others about opinion
