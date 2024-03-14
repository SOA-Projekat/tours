package main

import (
	"database-example/handler"
	"database-example/model"
	"database-example/repo"
	"database-example/service"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	connectionStr := "host=localhost user=postgres password=super dbname=gorm port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		print(err)
		return nil
	}

	database.AutoMigrate(&model.Student{})
	database.AutoMigrate(&model.Tour{})
	database.AutoMigrate(&model.Equipment{})
	database.AutoMigrate(&model.TourPoint{})
	return database
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}
	//student example
	studentRepo := &repo.StudentRepository{DatabaseConnection: database}
	studentService := &service.StudentService{StudentRepo: studentRepo}
	studentHandler := &handler.StudentHandler{StudentService: studentService}

	//equipment
	equipmentRepository := &repo.EquipmentRepository{DatabaseConnection: database}
	equipmentService := &service.EquipmentService{EquipmentRepo: equipmentRepository}
	equipmentHandler := &handler.EquipmentHandler{EquipmentService: equipmentService}

	//tour
	tourRepository := &repo.TourRepository{DatabaseConnection: database}
	tourService := &service.TourService{
		TourRepo:      tourRepository,
		EquipmentRepo: equipmentRepository,
	}
	tourHandler := &handler.TourHandler{TourService: tourService}

	//tourPoint

	tourPointRepository := &repo.TourPointRepository{DatabaseConnection: database}
	tourPointService := &service.TourPointService{TourPointRepo: tourPointRepository}
	tourPointHandler := &handler.TourPointHandler{TourPointService: tourPointService}

	router := mux.NewRouter().StrictSlash(true)

	//routes for student
	router.HandleFunc("/students/{id}", studentHandler.Get).Methods("GET")
	router.HandleFunc("/students", studentHandler.Create).Methods("POST")

	//routes for tours
	router.HandleFunc("/tours", tourHandler.Create).Methods("POST")
	router.HandleFunc("/tour/{id}", tourHandler.GetTourById).Methods("GET")
	router.HandleFunc("/tours/{userId}", tourHandler.GetToursForAuthor).Methods("GET")
	router.HandleFunc("/tours", tourHandler.UpdateTour).Methods("PUT")

	//routes for equipment
	router.HandleFunc("/equipments", equipmentHandler.Create).Methods("POST")
	router.HandleFunc("/equipment/{id}", equipmentHandler.GetEquipmentById).Methods("GET")
	router.HandleFunc("/equipments", equipmentHandler.UpdateEquipment).Methods("PUT")
	router.HandleFunc("/equipment/{id}", equipmentHandler.DeleteEquipment).Methods("DELETE")

	//routes for tour-equipment relations
	router.HandleFunc("/tours/{tourID}/equipments/{equipmentID}", tourHandler.AddEquipmentToTour).Methods("POST")

	//routes for tourPoints
	router.HandleFunc("/tourPoints", tourPointHandler.Create).Methods("POST")
	router.HandleFunc("/tours/{tourId}/points", tourPointHandler.GetTourPoints).Methods("GET")

	permitedHeaders := handlers.AllowedHeaders([]string{"Requested-With", "Content-Type", "Authorization"})
	permitedOrigins := handlers.AllowedOrigins([]string{"*"})
	permitedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8082", handlers.CORS(permitedHeaders, permitedOrigins, permitedMethods)(router)))
}
