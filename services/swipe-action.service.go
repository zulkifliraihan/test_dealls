package services

import (
	"fmt"
	"test_dealls/configs"
	dataStruct "test_dealls/data/struct"
	"test_dealls/models"
	"time"

	// "test_dealls/models"
	"test_dealls/repositories"
	"test_dealls/requests"

	"github.com/gin-gonic/gin"
)

type SwipeActionService interface {
	SwipeActionUser(ctx *gin.Context, data requests.SwipeActionRequest) ( dataStruct.FormatReturn, error)
}

type swipeActionService struct {
	userRepository repositories.UserRepository
	swipeRepository repositories.SwipeRepository
}

func NewSwipeActionService(
	userRepository repositories.UserRepository,
	swipeRepository repositories.SwipeRepository,
) SwipeActionService {
	return &swipeActionService{
		userRepository: userRepository,
		swipeRepository: swipeRepository,
	}
}

func (receiver *swipeActionService) SwipeActionUser(ctx *gin.Context, data requests.SwipeActionRequest) ( dataStruct.FormatReturn, error) {
    var formatReturn dataStruct.FormatReturn
    var findSwipedUser models.Swipe
	var currentUser models.User
	
	var directionSwipe string
	
	today := time.Now().Format("2006-01-02")
	fmt.Println("Today : ", today)
	
	currentUserRaw, _ := ctx.Get("currentUser")
	currentUser = currentUserRaw.(models.User)

	errFindSwipedUser := configs.DB.Where("swiping_user_id = ?", currentUser.ID).
			Where("swiped_user_id = ?", data.SwipedUserID).
			Where("DATE(created_at) = ?", today).
			Order("created_at DESC").
			Limit(1).
			First(&findSwipedUser).Error

	if errFindSwipedUser != nil {
		
	}
	fmt.Println("findSwipedUser : ", findSwipedUser)

	if findSwipedUser.ID != 0 {
		formatReturn = dataStruct.FormatReturn{
			Status:  false,
			Type:    "validation",
			Message: "",
			Data:    nil,
			Error:   "You've already swiped this user. Please choose someone else!",
		}

		return formatReturn, nil
	}
    userSwiped, _ := receiver.userRepository.FindById(data.SwipedUserID)

	if userSwiped == nil {
		formatReturn = dataStruct.FormatReturn{
			Status:  false,
			Type:    "validation",
			Message: "",
			Data:    nil,
			Error:   "User ID Not Found!",
		}

		return formatReturn, nil
	}


	if data.Direction == "right" {
		directionSwipe = "right"
	} else if data.Direction == "left" {
		directionSwipe = "left"
	} else {
		formatReturn = dataStruct.FormatReturn{
			Status:  false,
			Type:    "validation",
			Message: "",
			Data:    nil,
			Error:   "Format Direction Not Allowed!",
		}

		return formatReturn, nil
	}

	createDataSwipe := models.Swipe {
		SwipingUserID: currentUser.ID,
		SwipedUserID: uint(data.SwipedUserID),
		Direction: directionSwipe,
	}

	createdSwipe, errCreatedSwipe := receiver.swipeRepository.Save(&createDataSwipe)
	if errCreatedSwipe != nil {
		formatReturn = dataStruct.FormatReturn{
			Status:  false,
			Type:    "server",
			Message: "",
			Data:    nil,
			Error:   errCreatedSwipe.Error(),
		}
		return formatReturn, nil
	}

	configs.DB.First(&currentUser)
	currentUser.LastSwipe = time.Now()
	configs.DB.Save(&currentUser)

	dataReturn := map[string]interface{}{
		"userSwiped": userSwiped,
		"createdSwipe": createdSwipe,
	}

	formatReturn = dataStruct.FormatReturn{
		Status:  true,
		Type:    "created",
		Message: "",
		Data:    dataReturn,
		Error:   nil,
	}
	
    return formatReturn, nil
}
