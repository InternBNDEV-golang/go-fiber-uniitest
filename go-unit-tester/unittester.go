package unittester

import (
	"fmt"
	ds "go-fiber-unitest/domain/datasources"
	"go-fiber-unitest/domain/entities"
	repo "go-fiber-unitest/domain/repositories"
	sv "go-fiber-unitest/src/services"
	"log"

	"github.com/joho/godotenv"
)

//Import anything you need here

// init() will be call once before main()
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mongodb := ds.NewMongoDB(10)

	userMongo := repo.NewUsersRepository(mongodb)

	sv0 := sv.NewUsersService(userMongo)

	TestCaseGetAllUserStatusOK(sv0)
	TestCaseInsertNewAccountEmptyERROR(sv0)
	TestCaseInsertNewAccountStatusOK(sv0)
	TestCaseInsertNewAccountNotUserID(sv0)
	TestCaseInsertNewAccountNotEmail(sv0)

}
func TestCaseGetAllUserStatusOK(sv sv.IUsersService) {
	fmt.Printf("TestCaseGetAllUserStatusOK : ")
	input := []interface{}{}
	expected := []interface{}{[]entities.UserDataFormat{}, nil}
	checkOutputTypeOnly := true

	UnitTest(sv.GetAllUser, expected, input, checkOutputTypeOnly)
}
func TestCaseInsertNewAccountEmptyERROR(sv sv.IUsersService) {
	fmt.Println("TestCaseInsertNewAccountAccountEmptyERROR :  ")
	data := &entities.NewUserBody{
		Email:    "",
		UserID:   "",
		Username: "",
	}
	input := []interface{}{data}
	expected := []interface{}{false}
	checkOutputTypeOnly := false

	UnitTest(sv.InsertNewAccount, expected, input, checkOutputTypeOnly)
}

func TestCaseInsertNewAccountStatusOK(sv sv.IUsersService) {
	fmt.Println("TestCaseInsertNewAccountStatusOK :  ")
	data := &entities.NewUserBody{
		Email:    "test@test.com",
		UserID:   "test_1",
		Username: "ttest_1223",
	}
	input := []interface{}{data}
	expected := []interface{}{true}
	checkOutputTypeOnly := false

	UnitTest(sv.InsertNewAccount, expected, input, checkOutputTypeOnly)
}

func TestCaseInsertNewAccountNotUserID(sv sv.IUsersService) {
	fmt.Println("TestCaseInsertNewAccountNotUserID :  ")
	data := &entities.NewUserBody{
		Email:    "test@test.com",
		UserID:   "",
		Username: "ttest_1223",
	}
	input := []interface{}{data}
	expected := []interface{}{false}
	checkOutputTypeOnly := false

	UnitTest(sv.InsertNewAccount, expected, input, checkOutputTypeOnly)
}
func TestCaseInsertNewAccountNotEmail(sv sv.IUsersService) {
	fmt.Println("TestCaseInsertNewAccountNotEmail :  ")
	data := &entities.NewUserBody{
		Email:    "",
		UserID:   "test_1",
		Username: "ttest_1223",
	}
	input := []interface{}{data}
	expected := []interface{}{false}
	checkOutputTypeOnly := false

	UnitTest(sv.InsertNewAccount, expected, input, checkOutputTypeOnly)
}
