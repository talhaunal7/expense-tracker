package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/talhaunal7/expense-tracker/server/controller"
	"github.com/talhaunal7/expense-tracker/server/entity"
	"github.com/talhaunal7/expense-tracker/server/repository"
	"github.com/talhaunal7/expense-tracker/server/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var (
	server            *gin.Engine
	userRepository    repository.UserRepository
	userService       service.UserService
	userController    controller.UserController
	expenseRepository repository.ExpenseRepository
	expenseService    service.ExpenseService

	expenseController controller.ExpenseController
	db                *gorm.DB
	err               error
)

func init() {

	err = godotenv.Load()
	// change host name on env file to "localhost"
	if err != nil {
		log.Fatal(err.Error())
	}

	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	fmt.Println("values", dbHost, dbPort, dbUser, dbPassword, dbName)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	fmt.Println("dsn", dsn)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.AutoMigrate(&entity.User{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository = repository.NewUserRepository(db)
	userService = service.NewUserService(userRepository)
	userController = controller.NewUserController(userService)

	server = gin.Default()
}

func main() {

	basepath := server.Group("/v1")
	userController.RegisterUserRoutes(basepath)

	log.Fatal(server.Run(":8080"))
}
