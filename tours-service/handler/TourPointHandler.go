package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"fmt"
	"strconv"

	//"fmt"
	//"log"
	"net/http"
	//"strconv"
	"github.com/gorilla/mux"
)

type TourPointHandler struct {
	TourPointService *service.TourPointService
}

func (handler *TourPointHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var tourPoint model.TourPoint
	err := json.NewDecoder(req.Body).Decode(&tourPoint)
	if err != nil {
		println("error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.TourPointService.Create(&tourPoint)
	if err != nil {
		println("error while creating tour")
		writer.WriteHeader(http.StatusExpectationFailed)
		return

	}
	writer.WriteHeader((http.StatusCreated))
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(tourPoint)
}

func (handler *TourPointHandler) GetTourPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tourIdStr := vars["tourId"]
	tourId, err := strconv.Atoi(tourIdStr)
	if err != nil {
		http.Error(w, "Invalid tour ID", http.StatusBadRequest)
		return
	}

	tourPoints, err := handler.TourPointService.GetTourPointsByTourId(tourId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching tour points: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tourPoints)
}
