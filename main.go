package main

import (
	"go-fiber-unitest/configuration"
	ds "go-fiber-unitest/domain/datasources"
	repo "go-fiber-unitest/domain/repositories"
	gw "go-fiber-unitest/src/gateways"
	"go-fiber-unitest/src/middlewares"
	sv "go-fiber-unitest/src/services"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	// // // remove this before deploy ###################
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	// /// ############################################

	app := fiber.New(configuration.NewFiberConfiguration())
	middlewares.Logger(app)
	app.Use(recover.New())
	app.Use(cors.New())

	mongodb := ds.NewMongoDB(10)

	userMongo := repo.NewUsersRepository(mongodb)

	sv0 := sv.NewUsersService(userMongo)

	gw.NewHTTPGateway(app, sv0)

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	app.Listen(":" + PORT)
}
