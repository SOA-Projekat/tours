package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type EquipmentHandler struct {
	EquipmentService *service.EquipmentService
}

func (handler *EquipmentHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var equipment model.Equipment
	err := json.NewDecoder(req.Body).Decode(&equipment)
	if err != nil {
		println("error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.EquipmentService.CreateEquipment(&equipment)
	if err != nil {
		println("error while creting equipment")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Tpe", "application/json")
}

func (handler *EquipmentHandler) GetEquipmentById(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	equipment, err := handler.EquipmentService.GetEquipmentById(id)
	fmt.Println(equipment)
	writer.Header().Set("Content_Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(equipment)
}

func (handler *EquipmentHandler) UpdateEquipment(writer http.ResponseWriter, req *http.Request) {
	var equipment model.Equipment

	err := json.NewDecoder(req.Body).Decode(&equipment)
	if err != nil {
		println("error occured whiule parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.EquipmentService.UpdateEquipment(&equipment)
	if err != nil {
		println("error occured while updating tour")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(equipment)
}

func (handler *EquipmentHandler) DeleteEquipment(writer http.ResponseWriter, req *http.Request) {
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

	err = handler.EquipmentService.DeleteEquipment(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusNoContent) // 204 No Content is often used for successful deletes
}
