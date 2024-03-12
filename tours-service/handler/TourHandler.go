package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"fmt"
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

func (handler *TourHandler) GetToursForAuthor(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["userId"]
	converterId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("integer can't nbe converted to integer")
	}
	tours, err := handler.TourService.GetToursForAuthor(converterId)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(tours)
}

//TODO: probably keypoint logic is added here i suppose. ask others about opinion
