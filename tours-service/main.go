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
	//database.Exec("INSERT IGNORE INTO students VALUES ('aec7e123-233d-4a09-a289-75308ea5b7e6', 'Marko Markovic', 'Graficki dizajn')")
	return database
}

/*
	func startServer(handler *handler.StudentHandler) {
		router := mux.NewRouter().StrictSlash(true)

		router.HandleFunc("/students/{id}", handler.Get).Methods("GET")
		router.HandleFunc("/students", handler.Create).Methods("POST")

		permitedHeaders := handlers.AllowedHeaders([]string{"Requested-With", "Content-Type", "Authorization"})
		permitedOrigins := handlers.AllowedOrigins([]string{"*"})
		permitedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

		router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
		println("Server starting")
		log.Fatal(http.ListenAndServe(":8082", handlers.CORS(permitedHeaders, permitedOrigins, permitedMethods)(router)))

}
*/
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

	//tour
	tourRepository := &repo.TourRepository{DatabaseConnection: database}
	tourService := &service.TourService{TourRepo: tourRepository}
	tourHandler := &handler.TourHandler{TourService: tourService}

	router := mux.NewRouter().StrictSlash(true)

	//routes for student
	router.HandleFunc("/students/{id}", studentHandler.Get).Methods("GET")
	router.HandleFunc("/students", studentHandler.Create).Methods("POST")

	//routes for tours
	router.HandleFunc("/tours", tourHandler.Create).Methods("POST")
	router.HandleFunc("/tour/{id}", tourHandler.GetTourById).Methods("GET")
	router.HandleFunc("/tours/{userId}", tourHandler.GetToursForAuthor).Methods("GET")

	permitedHeaders := handlers.AllowedHeaders([]string{"Requested-With", "Content-Type", "Authorization"})
	permitedOrigins := handlers.AllowedOrigins([]string{"*"})
	permitedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8082", handlers.CORS(permitedHeaders, permitedOrigins, permitedMethods)(router)))
}
