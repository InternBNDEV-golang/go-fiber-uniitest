package services

import (
	"go-fiber-unitest/domain/entities"
	"go-fiber-unitest/domain/repositories"
)

type usersService struct {
	UsersRepository repositories.IUsersRepository
}

type IUsersService interface {
	GetAllUser() ([]entities.UserDataFormat, error)
	InsertNewAccount(data *entities.NewUserBody) bool
}

func NewUsersService(repo0 repositories.IUsersRepository) IUsersService {
	return &usersService{
		UsersRepository: repo0,
	}
}

func (sv usersService) GetAllUser() ([]entities.UserDataFormat, error) {
	userData, err := sv.UsersRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return userData, nil

}

func (sv usersService) InsertNewAccount(data *entities.NewUserBody) bool {
	if data.Email == "" && data.UserID == "" && data.Username == "" {
		return false
	} else if data.UserID == "" {
		return false
	} else if data.Email == "" {
		return false
	} else {
		return sv.UsersRepository.InsertNewUser(data)
	}
}
