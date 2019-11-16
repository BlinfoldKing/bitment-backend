package main

import (
	"fmt"
	"os"

	"github.com/ehardi19/bitment-backend/authentication"
	"github.com/ehardi19/bitment-backend/domain"
	"github.com/ehardi19/bitment-backend/repository"
	"github.com/ehardi19/bitment-backend/services"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASSWORD")

	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass))
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	orm := repository.NewRepository(db)
	auth := authentication.NewAuthentication()
	services := services.NewServices(orm, auth)

	services.CreateUser(domain.User{
		Name:     "root",
		Email:    "root",
		Password: "root",
		PIN:      "123456",
	})
}
