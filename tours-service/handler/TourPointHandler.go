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

func (handler *TourPointHandler) UpdateTourPoint(writer http.ResponseWriter, req *http.Request) {
	var tourPoint model.TourPoint

	err := json.NewDecoder(req.Body).Decode(&tourPoint)
	if err != nil {
		println("error occured whiule parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.TourPointService.UpdateTourPoint(&tourPoint)
	if err != nil {
		println("error occured while updating tour")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(tourPoint)
}

func (handler *TourPointHandler) DeleteTourPoint(writer http.ResponseWriter, req *http.Request) {
	// Extract ID from the request, assuming you're using gorilla/mux or a similar router
	vars := mux.Vars(req)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(writer, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(writer, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = handler.TourPointService.DeleteTourPoint(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusNoContent) // 204 No Content is often used for successful deletes
}
