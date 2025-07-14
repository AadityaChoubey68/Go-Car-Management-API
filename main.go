package main

import (
	"database/sql"

	"log"
	"net/http"
	"os"

	middleware "github.com/AadityaChoubey68/Go-Car-Management-API/Middleware"
	"github.com/AadityaChoubey68/Go-Car-Management-API/driver"
	carHandler "github.com/AadityaChoubey68/Go-Car-Management-API/handler/car"
	engineHandler "github.com/AadityaChoubey68/Go-Car-Management-API/handler/engine"
	loginHandler "github.com/AadityaChoubey68/Go-Car-Management-API/handler/login"
	carService "github.com/AadityaChoubey68/Go-Car-Management-API/service/car"
	engineService "github.com/AadityaChoubey68/Go-Car-Management-API/service/engine"
	carStore "github.com/AadityaChoubey68/Go-Car-Management-API/store/car"
	engineStore "github.com/AadityaChoubey68/Go-Car-Management-API/store/engine"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading evn file")
	}
	driver.InitDB()
	defer driver.CloseDB()

	db := driver.GetDB()
	carStore := carStore.New(db)
	CarService := carService.NewCarService(carStore)

	EngineStore := engineStore.New(db)
	EngineService := engineService.NewEngineService(EngineStore)

	carHandler := carHandler.NewCarHandler(CarService)
	engineHandler := engineHandler.NewEngineHandler(EngineService)

	router := mux.NewRouter()

	schemaFile := "store/schema.sql"
	if err := ExecuteSchemaFile(db, schemaFile); err != nil {
		log.Fatal("Error whlie executing the schema file : ", err)
	}

	router.HandleFunc("/login", loginHandler.LoginHandler).Methods("POST")

	//Middleware
	protected := router.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	protected.HandleFunc("/cars/{id}", carHandler.GetCarById).Methods("GET")
	protected.HandleFunc("/cars", carHandler.GetCarByBrand).Methods("GET")
	protected.HandleFunc("/cars", carHandler.CreateCar).Methods("POST")
	protected.HandleFunc("/cars/{id}", carHandler.UpdateCar).Methods("PUT")
	protected.HandleFunc("/cars/{id}", carHandler.DeleteCar).Methods("DELETE")

	protected.HandleFunc("/engine/{id}", engineHandler.GetEngById).Methods("GET")
	protected.HandleFunc("/engine", engineHandler.CreateEngine).Methods("POST")
	protected.HandleFunc("/engine/{id}", engineHandler.UpdateEngine).Methods("PUT")
	protected.HandleFunc("/engine/{id}", engineHandler.DeleteEngine).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// addr := fmt.Sprintf("%s", port)
	log.Printf("Server Listening port localhost : 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func ExecuteSchemaFile(db *sql.DB, fileName string) error {
	sqlFile, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	_, err = db.Exec(string(sqlFile))
	if err != nil {
		return err
	}
	return nil
}
