package services

import (
	"os"
	"strconv"
	dataStruct "test_dealls/data/struct"
	"test_dealls/repositories"
	"test_dealls/requests"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginService interface {
	LoginUser(data requests.LoginRequest) (dataStruct.FormatReturn, error)
}

type loginService struct {
	userRepository repositories.UserRepository
}

func NewLoginService(userRepository repositories.UserRepository) LoginService {
	return &loginService{
		userRepository: userRepository,
	}
}

func (receiver *loginService) LoginUser(data requests.LoginRequest) (dataStruct.FormatReturn, error) {
	var dataReturn dataStruct.FormatReturn

	findUser, _ := receiver.userRepository.FindByEmail(data.Email)

	if findUser == nil {
		dataReturn = dataStruct.FormatReturn{
			Status:  false,
			Type:    "validation-auth",
			Message: "Email address not found.",
			Data:    nil,
			Error:   nil,
		}
	} else {

		checkingPassword := bcrypt.CompareHashAndPassword([]byte(findUser[0].Password), []byte(data.Password))

		if checkingPassword != nil {
			dataReturn = dataStruct.FormatReturn{
				Status:  false,
				Type:    "validation-auth",
				Message: "Email or password is incorrect.",
				Data:    nil,
				Error:   nil,
			}
		} else {
			jwtExpiredHour := os.Getenv("JWT_EXPIRED_HOUR")
			expiredHour, _ := strconv.Atoi(jwtExpiredHour)
			secretJWT := os.Getenv("JWT_SECRET_KEY")
			expiredToken := time.Now().Add(time.Hour * time.Duration(expiredHour))

			claims := jwt.MapClaims{
				"expired":    expiredToken.Unix(),
				"expired_at": expiredToken,
				"user":       findUser,
				"userId":     findUser[0].ID,
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

			signedToken, _ := token.SignedString([]byte(secretJWT))

			valueData := map[string]interface{}{
				"users":         findUser,
				"token":         signedToken,
				"expired_token": expiredToken,
			}

			dataReturn = dataStruct.FormatReturn{
				Status:  true,
				Type:    "created",
				Message: "",
				Data:    valueData,
				Error:   nil,
			}
		}

	}

	return dataReturn, nil
}
