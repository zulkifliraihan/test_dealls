package handlers

import (
	"fmt"
	"test_dealls/helpers"
	"test_dealls/requests"
	authService "test_dealls/services/auth"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type LoginHandler interface {
	Login(ctx *gin.Context)
}

type loginHandler struct {
	loginService authService.LoginService
}

func NewLoginHandler(loginService authService.LoginService) LoginHandler {
	return &loginHandler{
		loginService: loginService,
	}
}

func (receiver *loginHandler) Login (ctx *gin.Context) {
	
	data := new(requests.LoginRequest)
	var formatResponse helpers.TypeReturnResponse

	err := ctx.ShouldBindJSON(&data)
	if err != nil {
        if validationErr, ok := err.(validator.ValidationErrors); ok {
            validation := helpers.ValidationCustomMessage(validationErr)
			fmt.Println(validation)
			formatResponse = helpers.ReturnResponse(false, "validation", nil, "", validation)

		} else {
			formatResponse = helpers.ReturnResponse(false, "server", nil, "", nil)

		}

		ctx.JSON(formatResponse.Code, formatResponse)
		return
	}

	users, _ := receiver.loginService.LoginUser(*data)
	
	formatResponse = helpers.ReturnResponse(users.Status, users.Type, users.Data, users.Message, users.Error)

	ctx.JSON(formatResponse.Code, formatResponse)
	return
}