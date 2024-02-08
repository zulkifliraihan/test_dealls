package middlewares

import (
	"fmt"
	"test_dealls/configs"
	"test_dealls/helpers"
	"test_dealls/models"
	"time"

	"github.com/gin-gonic/gin"
)

func UserPackageMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var formatResponse helpers.TypeReturnResponse
		var currentUser models.User
		var swipes []models.Swipe
		var swipesId []int

	
		today := time.Now().Format("2006-01-02")
		fmt.Println("Today In Middleware: ", today)
		
		currentUserRaw, _ := ctx.Get("currentUser")
		currentUser = currentUserRaw.(models.User)
	
		lastSwipe := currentUser.LastSwipe.Format("2006-01-02")
	
		errSwipes := configs.DB.Model(&swipes).Where("swiping_user_id = ?", currentUser.ID).Where("DATE(created_at) = ?", today).Pluck("swiped_user_id", &swipesId).Error
	
		if errSwipes != nil {
			formatResponse = helpers.ReturnResponse(false, "server", nil, "", errSwipes.Error())

			ctx.JSON(formatResponse.Code, formatResponse)
			ctx.Abort()
			return
		}
		
		
		if  !currentUser.Premium &&
			lastSwipe == today &&
			len(swipesId) >= 10 {
	
			fmt.Println("Middleware : lastSwipe lebih kecil dari today")
	
			formatResponse = helpers.ReturnResponse(false, "validation", nil, "", "You have reached the daily swipe limit. Upgrade to Premium for unlimited swipes.")

			ctx.JSON(formatResponse.Code, formatResponse)
			ctx.Abort()
			return
		}

		ctx.Next()
		return

	}
}