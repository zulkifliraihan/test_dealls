package middlewares

import (
	"fmt"
	"os"
	"strings"
	"test_dealls/configs"

	// dataStruct "test_dealls/data/struct"
	"test_dealls/helpers"
	"test_dealls/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTValidate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var formatResponse helpers.TypeReturnResponse

		headerAuth := ctx.GetHeader("Authorization")

		if headerAuth == "" {
			formatResponse = helpers.ReturnResponse(false, "middleware", nil, "Authorization header is required", nil)

			ctx.JSON(formatResponse.Code, formatResponse)
			ctx.Abort()
			return
		}

		// Extract Token
		splitAuth := strings.Split(headerAuth, " ")

		headerToken := splitAuth[1]
		headerAuthType := splitAuth[0]
		if headerAuthType != "Bearer" {
			formatResponse = helpers.ReturnResponse(false, "middleware", nil, "Type Authorization must be Bearer Token!", nil)

			ctx.JSON(formatResponse.Code, formatResponse)
			ctx.Abort()
			return
		}

		checkingToken, err := jwt.Parse(headerToken, jwtKeyFunc)

		if err != nil {
			formatResponse = helpers.ReturnResponse(false, "server", nil, "", err.Error())

			ctx.JSON(formatResponse.Code, formatResponse)
			ctx.Abort()
			return
		}

		// pathUrl := ctx.Path()
		// splitPath := strings.Split(pathUrl, "/")
		// roleInPath := splitPath[3]

		// // Checking Role Middleware
		dataToken := checkingToken.Claims.(jwt.MapClaims)
		fmt.Println("dataToken", dataToken["expired_at"])

		expiredToken := dataToken["expired_at"]
		currentTime := time.Now()
		compareTime, err := time.Parse(time.RFC3339, expiredToken.(string))

		if !compareTime.After(currentTime) {
			formatResponse = helpers.ReturnResponse(false, "middleware", nil, "Token has been expired!", nil)

			ctx.JSON(formatResponse.Code, formatResponse)
			ctx.Abort()
			return
		}

		// var dataCurrentUser []string

		// dataCurrentUser := []string(
		// 	Name : dataToken["user"].Name,
		// )

		var users []models.User
		errorDb := configs.DB.Where("id = ?", dataToken["userId"]).First(&users).Error
		if errorDb != nil {
			formatResponse = helpers.ReturnResponse(false, "server", nil, "", errorDb.Error())

			ctx.JSON(formatResponse.Code, formatResponse)
			ctx.Abort()
			return
		}
		fmt.Println("Users: ", users)

		ctx.Set("currentUser", users[0])

		ctx.Next()
		return
	}
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET_KEY")), nil
}
