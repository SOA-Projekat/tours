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

	//database.Exec("INSERT IGNORE INTO students VALUES ('aec7e123-233d-4a09-a289-75308ea5b7e6', 'Marko Markovic', 'Graficki dizajn')")
	return database
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	// Repositories
	studentRepo := &repo.StudentRepository{DatabaseConnection: database}
	tourRepository := &repo.TourRepository{DatabaseConnection: database}
	equipmentRepository := &repo.EquipmentRepository{DatabaseConnection: database}

	// Services
	studentService := &service.StudentService{StudentRepo: studentRepo}
	tourService := &service.TourService{
		TourRepo:      tourRepository,
		EquipmentRepo: equipmentRepository, // Pass equipment repository
	}
	equipmentService := &service.EquipmentService{EquipmentRepo: equipmentRepository}

	// Handlers
	studentHandler := &handler.StudentHandler{StudentService: studentService}
	tourHandler := &handler.TourHandler{TourService: tourService}
	equipmentHandler := &handler.EquipmentHandler{EquipmentService: equipmentService}

	// Router setup
	router := mux.NewRouter().StrictSlash(true)

	// Routes for students
	router.HandleFunc("/students/{id}", studentHandler.Get).Methods("GET")
	router.HandleFunc("/students", studentHandler.Create).Methods("POST")

	// Routes for tours
	router.HandleFunc("/tours", tourHandler.Create).Methods("POST")
	router.HandleFunc("/tour/{id}", tourHandler.GetTourById).Methods("GET")
	router.HandleFunc("/tours/{userId}", tourHandler.GetToursForAuthor).Methods("GET")
	router.HandleFunc("/tours", tourHandler.UpdateTour).Methods("PUT")
	router.HandleFunc("/tours/{id}/equipments", tourHandler.AddEquipmentToTour).Methods("POST")
	router.HandleFunc("/tour/{id}/equipment", tourHandler.GetEquipmentForTour).Methods("GET")

	// Routes for equipment
	router.HandleFunc("/equipments", equipmentHandler.Create).Methods("POST")
	router.HandleFunc("/equipment/{id}", equipmentHandler.GetEquipmentById).Methods("GET")
	router.HandleFunc("/equipments", equipmentHandler.UpdateEquipment).Methods("PUT")
	router.HandleFunc("/equipment/{id}", equipmentHandler.DeleteEquipment).Methods("DELETE")

	// CORS setup
	permittedHeaders := handlers.AllowedHeaders([]string{"Requested-With", "Content-Type", "Authorization"})
	permittedOrigins := handlers.AllowedOrigins([]string{"*"})
	permittedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

	// Start server
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8082", handlers.CORS(permittedHeaders, permittedOrigins, permittedMethods)(router)))
}
