package services

import (
	dataStruct "test_dealls/data/struct"
	"test_dealls/models"
	"test_dealls/repositories"
	"test_dealls/requests"

	"golang.org/x/crypto/bcrypt"
)

type RegisterService interface {
	RegisterUser(data requests.RegisterRequest) (dataStruct.FormatReturn, error)
}

type registerService struct {
	userRepository repositories.UserRepository
}

func NewRegisterService(userRepository repositories.UserRepository) RegisterService {
	return &registerService{
		userRepository: userRepository,
	}
}

func (receiver *registerService) RegisterUser(data requests.RegisterRequest) (dataStruct.FormatReturn, error) {
	var dataReturn dataStruct.FormatReturn

	findUser, _ := receiver.userRepository.FindByEmail(data.Email)

	if findUser != nil {
		dataReturn = dataStruct.FormatReturn{
			Status:  false,
			Type:    "validation-auth",
			Message: "This email address is already registered.",
			Data:    nil,
			Error:   nil,
		}
	} else {

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

		if err != nil {
			dataReturn = dataStruct.FormatReturn{
				Status:  false,
				Type:    "server",
				Message: "",
				Data:    nil,
				Error:   err.Error(),
			}
			return dataReturn, nil
		}

		user := models.User{
			Name:     data.Name,
			Email:    data.Email,
			Password: string(hashedPassword),
		}

		createdUser, err := receiver.userRepository.Save(&user)
		if err != nil {
			dataReturn = dataStruct.FormatReturn{
				Status:  false,
				Type:    "server",
				Message: "",
				Data:    nil,
				Error:   err.Error(),
			}
			return dataReturn, nil
		}

		dataReturn = dataStruct.FormatReturn{
			Status:  true,
			Type:    "created",
			Message: "",
			Data:    *createdUser,
			Error:   nil,
		}
	}

	return dataReturn, nil
}
