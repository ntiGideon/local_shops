package main

import (
	"baseProject/config"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Printf("Server starting on PORT %s \n", os.Getenv("PORT"))

	//db connections
	connectionString := fmt.Sprintf("host=%v port=%v user=%v dbname=%s password=%s sslmode=require", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASS"))
	_, err = config.ConnectDb(connectionString)
	if err != nil {
		log.Fatal(err)
	}

	//server here
	server := http.Server{
		Addr:           os.Getenv("PORT"),
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	serverError := server.ListenAndServe()
	if serverError != nil {
		log.Fatal("Error occurred :: ", serverError)
	}
}
