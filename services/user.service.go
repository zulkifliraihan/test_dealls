package services

import (
	"strconv"
	"test_dealls/models"
	"test_dealls/repositories"
)


type UserService interface {
	CreateUser(name, email string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) CreateUser(name, email string) (*models.User, error) {
	user := &models.User{Name: name, Email: email}
	return s.userRepo.Save(user)
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.FindAll()
}

func (s *userService) FindUser(typefind string, data string) ([]models.User, error) {
	if typefind == "email" {
		return s.userRepo.FindByEmail(data)
	} else if typefind == "id" {
		
		num, err := strconv.Atoi(data)
		if err != nil {
			return nil, err
		}

		return s.userRepo.FindById(num)
	}
	return s.userRepo.FindAll()
}
