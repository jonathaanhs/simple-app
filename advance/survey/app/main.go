package main

import (
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/learn/simple-app/advance/survey/delivery"
	"github.com/learn/simple-app/advance/survey/repository"
	"github.com/learn/simple-app/advance/survey/usecase"
	_ "github.com/lib/pq" //import for postgres driver
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")
	dbUser := os.Getenv("db_user")
	dbPass := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dsn := "postgresql://" + dbUser + ":" + dbPass + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=disable"
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Minute * 10)

	userRepo := repository.NewUserRepository(db)
	surveyUsecase := usecase.NewSurveyUsecase(userRepo)
	surveyHandler := delivery.NewSurveyHandler(surveyUsecase)

	surveyHandler.StartSurvey()
}
