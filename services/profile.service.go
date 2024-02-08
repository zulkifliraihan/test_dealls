package services

import (
	"fmt"
	"math/rand"
	"sort"
	"test_dealls/configs"
	dataStruct "test_dealls/data/struct"
	"test_dealls/models"
	"test_dealls/repositories"
	"time"

	"github.com/gin-gonic/gin"
)

type ProfileService interface {
	ProfileUser(ctx *gin.Context) ( dataStruct.FormatReturn, error)
}

type profileService struct {
	userRepository repositories.UserRepository
	swipeRepository repositories.SwipeRepository
}

func NewProfileService(
	userRepository repositories.UserRepository,
	swipeRepository repositories.SwipeRepository,
) ProfileService {
	return &profileService{
		userRepository: userRepository,
		swipeRepository: swipeRepository,
	}
}

func (receiver *profileService) ProfileUser(ctx *gin.Context) ( dataStruct.FormatReturn, error) {
    var dataReturn dataStruct.FormatReturn
    var users []models.User
    var user []models.User
    var swipesId []int
    var usersId []int
	
	today := time.Now().Format("2006-01-02")
	fmt.Println("Today : ", today)

	users, _ = receiver.userRepository.QueryWhere("id NOT IN ?", swipesId)

	errUser := configs.DB.Model(&users).Where("id NOT IN ?", swipesId).Pluck("id", &usersId).Error
	if errUser != nil {
		dataReturn = dataStruct.FormatReturn{
			Status:  false,
			Type:    "server",
			Message: "",
			Data:    nil,
			Error:   errUser.Error(),
		}

		return dataReturn, nil
	}

	sort.Ints(usersId)
	smallestUserId := usersId[0]
	largestUserId := usersId[len(usersId)-1]

	randomUserId := rand.Intn(largestUserId-smallestUserId+1) + smallestUserId

    user, _ = receiver.userRepository.FindById(randomUserId)

	data := map[string]interface{}{
		"user": user,
		"swiped_user_id_today": swipesId,
		"available_user_today": users,
	}

	dataReturn = dataStruct.FormatReturn{
		Status:  true,
		Type:    "get",
		Message: "",
		Data:    data,
		Error:   nil,
	}
	
    return dataReturn, nil
}
