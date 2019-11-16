package main

import (
	"fmt"
	"os"

	"github.com/ehardi19/bitment-backend/authentication"
	"github.com/ehardi19/bitment-backend/handler"
	"github.com/ehardi19/bitment-backend/repository"
	"github.com/ehardi19/bitment-backend/services"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
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

	h := handler.NewHandler(services)

	e := echo.New()

	e.GET("/", h.HelloWorld)
	e.POST("/login", h.Login)
	e.POST("/register", h.Register)

	user := e.Group("/user")
	user.GET("", h.GetAllUser)
	user.POST("", h.CreateUser)
	user.GET("/:id", h.GetUserByID)
	user.PUT("/:id", h.UpdateUser)
	user.PUT("/:id/password", h.ChangePassword)
	user.PUT("/:id/pin", h.ChangePIN)
	user.GET("/:id/balance", h.GetAccountBalanceByUser)
	user.GET("/:id/credit", h.GetAccountCreditByUser)

	goal := e.Group("/goal")
	goal.GET("/user/:user_id", h.GetAllGoalByUser)
	goal.POST("", h.CreateGoal)
	goal.GET("/:id", h.GetGoalByID)

	investment := e.Group("/investment")
	investment.GET("/user/:user_id", h.GetAllInvestmentByUser)
	investment.GET("/user/:user_id/total", h.GetTotalInvestmentProjectionByUser)
	investment.POST("", h.CreateInvestment)
	investment.GET("/:id", h.GetInvestmentByID)

	transaction := e.Group("/transaction")
	transaction.GET("/user/:user_id", h.GetAllTransactionByUser)
	transaction.POST("", h.CreateTransaction)
	transaction.GET("/:id", h.GetTransactionByID)

	item := e.Group("/item")
	item.GET("", h.GetAllItem)
	item.GET("/:id", h.GetItemByID)

	vendor := e.Group("/vendor")
	vendor.GET("", h.GetAllVendor)
	vendor.GET("/:id", h.GetVendorByID)

	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(port))
}
