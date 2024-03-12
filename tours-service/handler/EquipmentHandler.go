package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"

	"net/http"
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
