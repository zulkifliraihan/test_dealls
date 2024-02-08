package services

import (
	"test_dealls/configs"
	dataStruct "test_dealls/data/struct"
	"test_dealls/models"
	"test_dealls/repositories"

	"github.com/gin-gonic/gin"
)

type UpgradePackageService interface {
	UpgradePackageUser(ctx *gin.Context) ( dataStruct.FormatReturn, error)
}

type upgradePackageService struct {
	userRepository repositories.UserRepository
}

func NewUpgradePackageService(
	userRepository repositories.UserRepository,
	swipeRepository repositories.SwipeRepository,
) UpgradePackageService {
	return &upgradePackageService{
		userRepository: userRepository,
	}
}

func (receiver *upgradePackageService) UpgradePackageUser(ctx *gin.Context) ( dataStruct.FormatReturn, error) {
    var dataReturn dataStruct.FormatReturn
	var currentUser models.User

	currentUserRaw, _ := ctx.Get("currentUser")
	currentUser = currentUserRaw.(models.User)

	// INSERT INTEGRATION PAYMENT GATEWAY API IN HIRE //
	//                                                //
	// INSERT INTEGRATION PAYMENT GATEWAY API IN HIRE //

	configs.DB.First(&currentUser)
	currentUser.Premium = true
	configs.DB.Save(&currentUser)

	dataReturn = dataStruct.FormatReturn{
		Status:  true,
		Type:    "updated",
		Message: "",
		Data:    currentUser,
		Error:   nil,
	}
	
    return dataReturn, nil
}
