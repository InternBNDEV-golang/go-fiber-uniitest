package unittest

import (
	ds "go-fiber-unitest/domain/datasources"
	repo "go-fiber-unitest/domain/repositories"
	sv "go-fiber-unitest/src/services"
	"go-fiber-unitest/unittest/test"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env.test")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mongodb := ds.NewMongoDB(10)

	userMongo := repo.NewUsersRepository(mongodb)

	sv0 := sv.NewUsersService(userMongo)

	test.RunTestUserService(sv0)

}
